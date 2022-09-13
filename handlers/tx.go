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
	"github.com/kaifei-bianjie/msg-parser/modules/mt"
	msgsdktypes "github.com/kaifei-bianjie/msg-parser/types"
	types2 "github.com/tendermint/tendermint/abci/types"
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

	blockResults, err := client.BlockResults(context.Background(), &b)
	if err != nil {
		time.Sleep(1 * time.Second)
		blockResults, err = client.BlockResults(context.Background(), &b)
		if err != nil {
			return &blockDoc, nil, utils.ConvertErr(b, "", "ParseBlockResult", err)
		}
	}

	if len(block.Block.Txs) != len(blockResults.TxsResults) {
		return nil, nil, utils.ConvertErr(b, "", "block.Txs length not equal blockResult", nil)
	}

	txDocs := make([]*models.Tx, 0, len(block.Block.Txs))
	if len(block.Block.Txs) > 0 {
		for i, v := range block.Block.Txs {
			txResult := blockResults.TxsResults[i]
			txDoc, err := parseTx(v, txResult, block.Block, i)
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

func parseTx(txBytes types.Tx, txResult *types2.ResponseDeliverTx, block *types.Block, index int) (models.Tx, error) {
	var (
		docTx          models.Tx
		docTxMsgs      []msgsdktypes.TxMsg
		includeCfgType bool
	)
	txHash := utils.BuildHex(txBytes.Hash())

	docTx.Time = block.Time.Unix()
	docTx.Height = block.Height
	docTx.TxHash = txHash
	docTx.Status = parseTxStatus(txResult.Code)
	if docTx.Status == constant.TxStatusFail {
		docTx.Log = txResult.Log
	}

	docTx.EventsNew = parseABCILogs(txResult.Log)
	docTx.TxIndex = uint32(index)
	docTx.TxId = block.Height*100000 + int64(index)

	eventsIndexMap := make(map[uint32]models.EventNew)
	if txResult.Code == 0 {
		eventsIndexMap = splitEvents(txResult.Log)
	}

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("errTag", "TxDecoder"),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return docTx, nil
	}
	docTx.GasUsed = txResult.GasUsed
	docTx.Fee = msgsdktypes.BuildFee(authTx.GetFee(), authTx.GetGas())
	docTx.Memo = authTx.GetMemo()

	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return docTx, nil
	}

	for i, v := range msgs {
		msgDocInfo := _parser.HandleTxMsg(v)
		if len(msgDocInfo.Addrs) == 0 {
			if len(_filterMap) > 0 {
				if len(docTx.EventsNew) > i {
					docTx.EventsNew[i].Events = []models.Event{}
				}
				//add empty msg for msgIndex match
				docTxMsgs = append(docTxMsgs, msgsdktypes.TxMsg{Type: "no setting type"})
			}
			continue
		}
		if len(_filterMap) > 0 {
			_, ok := _filterMap[msgDocInfo.DocTxMsg.Type]
			if ok && !includeCfgType {
				includeCfgType = true
			}
			if !ok {
				if len(docTx.EventsNew) > i {
					docTx.EventsNew[i].Events = []models.Event{}
				}
				docTxMsgs = append(docTxMsgs, msgsdktypes.TxMsg{Type: msgDocInfo.DocTxMsg.Type})
				continue
			}
		}

		switch msgDocInfo.DocTxMsg.Type {
		case MsgTypeMTIssueDenom:
			if docTx.Status == constant.TxStatusFail {
				break
			}

			// get denom_id from events then set to msg, because this msg hasn't denom_id
			denomId := ParseAttrValueFromEvents(docTx.EventsNew[i].Events, EventTypeIssueDenom, AttrKeyDenomId)
			msgDocInfo.DocTxMsg.Msg.(*mt.DocMsgMTIssueDenom).Id = denomId
		case MsgTypeMintMT:
			if docTx.Status == constant.TxStatusFail {
				break
			}

			// get mt_id from events then set to msg, because this msg hasn't mt_id
			mtId := ParseAttrValueFromEvents(docTx.EventsNew[i].Events, EventTypeMintMT, AttrKeyMTId)
			msgDocInfo.DocTxMsg.Msg.(*mt.DocMsgMTMint).Id = mtId
		case MsgTypeIBCTransfer:
			if ibcTranferMsg, ok := msgDocInfo.DocTxMsg.Msg.(*ibc.DocMsgTransfer); ok {
				if val, exist := eventsIndexMap[uint32(i)]; exist {
					ibcTranferMsg.PacketId = buildPacketId(val.Events)
					msgDocInfo.DocTxMsg.Msg = ibcTranferMsg
				}
				if _conf.Server.IgnoreIbcHeader {
					for id, one := range docTx.EventsNew {
						if one.MsgIndex == uint32(i) {
							docTx.EventsNew[id].Events = hookEvents(docTx.EventsNew[id].Events, removePacketDataHexOfIbcTxEvents)
						}
					}
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
			if _conf.Server.IgnoreIbcHeader {
				timeOutMsg, ok := msgDocInfo.DocTxMsg.Msg.(*ibc.DocMsgRecvPacket)
				if ok {
					timeOutMsg.ProofCommitment = "ignore ibc ProofCommitment info"
					msgDocInfo.DocTxMsg.Msg = timeOutMsg
				}
				for id, one := range docTx.EventsNew {
					if one.MsgIndex == uint32(i) {
						docTx.EventsNew[id].Events = hookEvents(docTx.EventsNew[id].Events, removePacketDataHexOfIbcTxEvents)
					}
				}
			}
		case MsgTypeUpdateClient:
			if _conf.Server.IgnoreIbcHeader {
				updateClientMsg, ok := msgDocInfo.DocTxMsg.Msg.(*ibc.DocMsgUpdateClient)
				if ok {
					updateClientMsg.Header = "ignore ibc header info"
					msgDocInfo.DocTxMsg.Msg = updateClientMsg
				}
				for id, one := range docTx.EventsNew {
					if one.MsgIndex == uint32(i) {
						docTx.EventsNew[id].Events = hookEvents(docTx.EventsNew[id].Events, removeHeaderOfUpdateClientEvents)
					}
				}
			}
		case MsgTypeTimeout:
			if _conf.Server.IgnoreIbcHeader {
				timeOutMsg, ok := msgDocInfo.DocTxMsg.Msg.(*ibc.DocMsgTimeout)
				if ok {
					timeOutMsg.ProofUnreceived = "ignore ibc ProofUnreceived info"
					msgDocInfo.DocTxMsg.Msg = timeOutMsg
				}
			}
		}
		if i == 0 {
			docTx.Type = msgDocInfo.DocTxMsg.Type
		}
		if docTx.Type == "" {
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

const (
	EventTypeIssueDenom = "issue_denom"
	EventTypeMintMT     = "mint_mt"
	AttrKeyDenomId      = "denom_id"
	AttrKeyMTId         = "mt_id"
)

func ParseAttrValueFromEvents(events []models.Event, typ, attrKey string) string {
	for _, val := range events {
		if val.Type == typ {
			for _, attr := range val.Attributes {
				if attr.Key == attrKey {
					return attr.Value
				}
			}
		}
	}
	return ""
}
