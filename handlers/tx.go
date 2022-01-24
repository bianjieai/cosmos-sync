package handlers

import (
	"fmt"
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/msgparser"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	"github.com/kaifei-bianjie/msg-parser/codec"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/modules/ibc"
	msgsdktypes "github.com/kaifei-bianjie/msg-parser/types"
	aTypes "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"golang.org/x/net/context"
	"strings"
	"time"
)

var (
	_parser    msgparser.MsgParser
	_conf      *config.Config
	_filterMap map[string]string
)

func InitRouter(conf *config.Config) {
	_conf = conf
	var router msgparser.Router
	if conf.Server.SupportModules != "" {
		modules := strings.Split(conf.Server.SupportModules, ",")
		msgRoute := msgparser.NewRouter()
		for _, one := range modules {
			fn, exist := msgparser.RouteHandlerMap[one]
			if !exist {
				logger.Fatal("no support module: " + one)
			}
			msgRoute = msgRoute.AddRoute(one, fn)
			switch one {
			case msgparser.IbcRouteKey:
				msgRoute = msgRoute.AddRoute(msgparser.IbcTransferRouteKey, msgparser.RouteHandlerMap[one])
			case msgparser.TIbcRouteKey:
				msgRoute = msgRoute.AddRoute(msgparser.TIbcTransferRouteKey, msgparser.RouteHandlerMap[one])
			}
		}
		if msgRoute.GetRoutesLen() > 0 {
			router = msgRoute
		}
	} else {
		router = msgparser.RegisteRouter()
	}

	// check and remove disable support module route path
	if conf.Server.DenyModules != "" {
		modules := strings.Split(conf.Server.DenyModules, ",")
		for _, one := range modules {
			_, exist := msgparser.RouteHandlerMap[one]
			if !exist {
				logger.Fatal("disable no exist module: " + one)
			}
			if router.HasRoute(one) {
				switch one {
				case msgparser.IbcRouteKey:
					router.RemoveRoute(msgparser.IbcRouteKey)
					router.RemoveRoute(msgparser.IbcTransferRouteKey)
				case msgparser.TIbcRouteKey:
					router.RemoveRoute(msgparser.TIbcRouteKey)
					router.RemoveRoute(msgparser.TIbcTransferRouteKey)
				default:
					router.RemoveRoute(one)
				}
			}
		}
	}
	_parser = msgparser.NewMsgParser(router)

	if conf.Server.Bech32AccPrefix != "" {
		initBech32Prefix(conf.Server.Bech32AccPrefix)
	}
	//ibc-zone
	if filterMsgType := models.GetSrvConf().SupportTypes; filterMsgType != "" {
		msgTypes := strings.Split(filterMsgType, ",")
		_filterMap = make(map[string]string, len(msgTypes))
		for _, val := range msgTypes {
			_filterMap[val] = val
		}
	}
}

func ParseBlockAndTxs(b int64, client *pool.Client) (*models.Block, []*models.Tx, error) {
	var (
		blockDoc models.Block
		block    *ctypes.ResultBlock
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if v, err := client.Block(ctx, &b); err != nil {
		time.Sleep(1 * time.Second)
		if v2, err := client.Block(ctx, &b); err != nil {
			return &blockDoc, nil, utils.ConvertErr(b, "", "ParseBlock", err)
		} else {
			block = v2
		}
	} else {
		block = v
	}
	blockDoc = models.Block{
		Height:   block.Block.Height,
		Time:     block.Block.Time.Unix(),
		Hash:     block.Block.Header.Hash().String(),
		Txn:      int64(len(block.Block.Data.Txs)),
		Proposer: block.Block.ProposerAddress.String(),
	}

	txResultMap := handleTxResult(client, block.Block)

	txDocs := make([]*models.Tx, 0, len(block.Block.Txs))
	if len(block.Block.Txs) > 0 {
		for index, v := range block.Block.Txs {
			txHash := utils.BuildHex(v.Hash())
			txResult, ok := txResultMap[txHash]
			if !ok || txResult == nil {
				return &blockDoc, txDocs, utils.ConvertErr(block.Block.Height, txHash, "TxResult",
					fmt.Errorf("no found"))
			}
			txDoc, err := parseTx(v, txResult, block.Block, index)
			if err != nil {
				return &blockDoc, txDocs, err
			}
			if txDoc.TxHash != "" && len(txDoc.Type) > 0 {
				txDocs = append(txDocs, &txDoc)
			}
		}
	}

	return &blockDoc, txDocs, nil
}

func parseTx(txBytes types.Tx, txResult *ctypes.ResultTx, block *types.Block, index int) (models.Tx, error) {
	var (
		docTx          models.Tx
		docTxMsgs      []msgsdktypes.TxMsg
		includeCfgType bool
	)
	txHash := utils.BuildHex(txBytes.Hash())

	docTx.Time = block.Time.Unix()
	docTx.Height = block.Height
	docTx.TxHash = txHash
	docTx.Status = parseTxStatus(txResult.TxResult.Code)
	if docTx.Status == constant.TxStatusFail {
		docTx.Log = txResult.TxResult.Log
	}

	docTx.Events = parseEvents(txResult.TxResult.Events)
	docTx.EventsNew = parseABCILogs(txResult.TxResult.Log)
	docTx.TxIndex = uint32(index)
	eventsIndexMap := make(map[uint32]models.EventNew)
	if txResult.TxResult.Code == 0 {
		eventsIndexMap = splitEvents(txResult.TxResult.Log)
	}

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("errTag", "TxDecoder"),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return docTx, nil
	}
	docTx.Fee = msgsdktypes.BuildFee(authTx.GetFee(), authTx.GetGas())
	docTx.Memo = authTx.GetMemo()

	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return docTx, nil
	}

	for i, v := range msgs {
		msgDocInfo := _parser.HandleTxMsg(v)
		if len(msgDocInfo.Addrs) == 0 {
			continue
		}
		if len(_filterMap) > 0 {
			if _, ok := _filterMap[msgDocInfo.DocTxMsg.Type]; ok && !includeCfgType {
				includeCfgType = true
			}
		}
		switch msgDocInfo.DocTxMsg.Type {
		case MsgTypeIBCTransfer:
			if ibcTranferMsg, ok := msgDocInfo.DocTxMsg.Msg.(*ibc.DocMsgTransfer); ok {
				if val, exist := eventsIndexMap[uint32(i)]; exist {
					ibcTranferMsg.PacketId = buildPacketId(val.Events)
					msgDocInfo.DocTxMsg.Msg = ibcTranferMsg
				}

			} else {
				logger.Warn("ibc transfer handler packet_id failed", logger.String("errTag", "TxMsg"),
					logger.String("txhash", txHash),
					logger.Int("msg_index", i),
					logger.Int64("height", block.Height))
			}
		case MsgTypeTIBCRecvPacket:
			//docTx.Events = updateEvents(docTx.Events, UnmarshalTibcAcknowledgement)
			for id, one := range docTx.EventsNew {
				if one.MsgIndex == uint32(i) {
					docTx.EventsNew[id].Events = updateEvents(docTx.EventsNew[id].Events, UnmarshalTibcAcknowledgement)
				}
			}
		case MsgTypeRecvPacket:
			//docTx.Events = updateEvents(docTx.Events, UnmarshalAcknowledgement)
			for id, one := range docTx.EventsNew {
				if one.MsgIndex == uint32(i) {
					docTx.EventsNew[id].Events = updateEvents(docTx.EventsNew[id].Events, UnmarshalAcknowledgement)
				}
			}

		}
		if i == 0 {
			docTx.Type = msgDocInfo.DocTxMsg.Type
		}

		docTx.Signers = append(docTx.Signers, removeDuplicatesFromSlice(msgDocInfo.Signers)...)
		docTx.Addrs = append(docTx.Addrs, removeDuplicatesFromSlice(msgDocInfo.Addrs)...)
		docTxMsgs = append(docTxMsgs, msgDocInfo.DocTxMsg)
		docTx.Types = append(docTx.Types, msgDocInfo.DocTxMsg.Type)
	}

	docTx.Addrs = removeDuplicatesFromSlice(docTx.Addrs)
	docTx.Types = removeDuplicatesFromSlice(docTx.Types)
	docTx.Signers = removeDuplicatesFromSlice(docTx.Signers)
	docTx.DocTxMsgs = docTxMsgs

	//setting type but not included in tx,skip this tx
	if len(_filterMap) > 0 && !includeCfgType {
		logger.Warn("skip tx for no include setting types",
			logger.String("types", strings.Join(docTx.Types, ",")),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return models.Tx{}, nil
	}

	// don't save txs which have not parsed
	if docTx.Type == "" {
		logger.Warn(constant.NoSupportMsgTypeTag,
			logger.String("errTag", "TxMsg"),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return models.Tx{}, nil
	}

	return docTx, nil
}

func buildPacketId(events []models.Event) string {
	if len(events) > 0 {
		var mapKeyValue map[string]string
		for _, e := range events {
			if len(e.Attributes) > 0 && e.Type == constant.IbcTransferEventTypeSendPacket {
				mapKeyValue = make(map[string]string, len(e.Attributes))
				for _, v := range e.Attributes {
					mapKeyValue[string(v.Key)] = string(v.Value)
				}
				break
			}
		}

		if len(mapKeyValue) > 0 {
			scPort := mapKeyValue[constant.IbcTransferEventAttriKeyPacketScPort]
			scChannel := mapKeyValue[constant.IbcTransferEventAttriKeyPacketScChannel]
			dcPort := mapKeyValue[constant.IbcTransferEventAttriKeyPacketDcPort]
			dcChannel := mapKeyValue[constant.IbcTransferEventAttriKeyPacketDcChannels]
			sequence := mapKeyValue[constant.IbcTransferEventAttriKeyPacketSequence]
			return fmt.Sprintf("%v%v%v%v%v", scPort, scChannel, dcPort, dcChannel, sequence)
		}
	}
	return ""
}

func parseTxStatus(code uint32) uint32 {
	if code == 0 {
		return constant.TxStatusSuccess
	} else {
		return constant.TxStatusFail
	}
}

func splitEvents(log string) map[uint32]models.EventNew {
	var eventDocs []models.EventNew
	if log != "" {
		utils.UnMarshalJsonIgnoreErr(log, &eventDocs)

	}

	msgIndexMap := make(map[uint32]models.EventNew, len(eventDocs))
	for _, val := range eventDocs {
		msgIndexMap[val.MsgIndex] = val
	}
	return msgIndexMap
}

func updateEvents(events []models.Event, fn func([]byte) string) []models.Event {

	for i, e := range events {
		if e.Type != constant.IbcRecvPacketEventTypeWriteAcknowledge {
			continue
		}
		if len(e.Attributes) > 0 {
			for j, v := range e.Attributes {
				if v.Key == constant.IbcRecvPacketEventAttriKeyPacketAck {
					attr := models.KvPair{
						Key:   string(v.Key),
						Value: string(v.Value),
					}
					attr.Value = fn([]byte(v.Value))
					e.Attributes[j] = attr
				}
			}
		}
		one := models.Event{
			Type:       e.Type,
			Attributes: e.Attributes,
		}
		events[i] = one
	}
	return events
}

func parseEvents(events []aTypes.Event) []models.Event {
	var eventDocs []models.Event
	if len(events) > 0 {
		for _, e := range events {
			var kvPairDocs []models.KvPair
			if len(e.Attributes) > 0 {
				for _, v := range e.Attributes {
					kvPairDocs = append(kvPairDocs, models.KvPair{
						Key:   string(v.Key),
						Value: string(v.Value),
					})
				}
			}
			eventDocs = append(eventDocs, models.Event{
				Type:       e.Type,
				Attributes: kvPairDocs,
			})
		}
	}

	return eventDocs
}

// parseABCILogs attempts to parse a stringified ABCI tx log into a slice of
// EventNe types. It ignore error upon JSON decoding failure.
func parseABCILogs(logs string) []models.EventNew {
	var res []models.EventNew
	utils.UnMarshalJsonIgnoreErr(logs, &res)
	return res
}

func removeDuplicatesFromSlice(data []string) (result []string) {
	tempSet := make(map[string]string, len(data))
	for _, val := range data {
		if _, ok := tempSet[val]; ok || val == "" {
			continue
		}
		tempSet[val] = val
	}
	for one := range tempSet {
		result = append(result, one)
	}
	return
}
