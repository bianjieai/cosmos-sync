package handlers

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/msgparser"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	"github.com/kaifei-bianjie/msg-parser/codec"
	msgsdktypes "github.com/kaifei-bianjie/msg-parser/types"
	aTypes "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/txn"
	"strings"
	"time"
)

var _parser msgparser.MsgParser

func InitRouter(conf *config.Config) {
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
}

func ParseBlockAndTxs(b int64, client *pool.Client) (*models.Block, []*models.Tx, []txn.Op, error) {
	var (
		blockDoc models.Block
		block    *ctypes.ResultBlock
		txnOps   []txn.Op
	)
	ctx := context.Background()
	if v, err := client.Block(ctx, &b); err != nil {
		time.Sleep(1 * time.Second)
		if v2, err := client.Block(ctx, &b); err != nil {
			return &blockDoc, nil, txnOps, utils.ConvertErr(b, "", "ParseBlock", err)
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

	txDocs := make([]*models.Tx, 0, len(block.Block.Txs))
	if len(block.Block.Txs) > 0 {
		for _, v := range block.Block.Txs {
			txDoc, ops, err := parseTx(client, v, block.Block)
			if err != nil {
				return &blockDoc, txDocs, txnOps, err
			}
			if txDoc.TxHash != "" && len(txDoc.Type) > 0 {
				txDocs = append(txDocs, &txDoc)
				if len(ops) > 0 {
					txnOps = append(txnOps, ops...)
				}
			}
		}
	}

	return &blockDoc, txDocs, txnOps, nil
}

func parseTx(c *pool.Client, txBytes types.Tx, block *types.Block) (models.Tx, []txn.Op, error) {
	var (
		docTx models.Tx

		docTxMsgs []msgsdktypes.TxMsg
		txnOps    []txn.Op
	)
	txHash := utils.BuildHex(txBytes.Hash())
	ctx := context.Background()
	txResult, err := c.Tx(ctx, txBytes.Hash(), false)
	if err != nil {
		time.Sleep(1 * time.Second)
		if v, err := c.Tx(ctx, txBytes.Hash(), false); err != nil {
			return docTx, txnOps, utils.ConvertErr(block.Height, txHash, "TxResult", err)
		} else {
			txResult = v
		}
	}
	docTx.Time = block.Time.Unix()
	docTx.Height = txResult.Height
	docTx.TxHash = txHash
	docTx.Status = parseTxStatus(txResult.TxResult.Code)
	if docTx.Status == constant.TxStatusFail {
		docTx.Log = txResult.TxResult.Log
	}

	docTx.Events = parseEvents(txResult.TxResult.Events)
	docTx.EventsNew = parseABCILogs(txResult.TxResult.Log)
	docTx.TxIndex = txResult.Index

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("errTag", "TxDecoder"),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return docTx, txnOps, nil
	}
	docTx.Fee = msgsdktypes.BuildFee(authTx.GetFee(), authTx.GetGas())
	docTx.Memo = authTx.GetMemo()

	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return docTx, txnOps, nil
	}

	for i, v := range msgs {
		msgDocInfo, ops := _parser.HandleTxMsg(v)
		if len(msgDocInfo.Addrs) == 0 {
			continue
		}
		if i == 0 {
			docTx.Type = msgDocInfo.DocTxMsg.Type
		}

		docTx.Signers = append(docTx.Signers, removeDuplicatesFromSlice(msgDocInfo.Signers)...)
		docTx.Addrs = append(docTx.Addrs, removeDuplicatesFromSlice(msgDocInfo.Addrs)...)
		docTxMsgs = append(docTxMsgs, msgDocInfo.DocTxMsg)
		docTx.Types = append(docTx.Types, msgDocInfo.DocTxMsg.Type)
		if len(ops) > 0 {
			txnOps = append(txnOps, ops...)
		}
	}

	docTx.Addrs = removeDuplicatesFromSlice(docTx.Addrs)
	docTx.Types = removeDuplicatesFromSlice(docTx.Types)
	docTx.Signers = removeDuplicatesFromSlice(docTx.Signers)
	docTx.DocTxMsgs = docTxMsgs

	// don't save txs which have not parsed
	if docTx.Type == "" {
		logger.Warn(constant.NoSupportMsgTypeTag,
			logger.String("errTag", "TxMsg"),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return models.Tx{}, txnOps, nil
	}

	return docTx, txnOps, nil
}

func parseTxStatus(code uint32) uint32 {
	if code == 0 {
		return constant.TxStatusSuccess
	} else {
		return constant.TxStatusFail
	}
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
