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
	"github.com/kaifei-bianjie/msg-parser/modules/evm"
	msgsdktypes "github.com/kaifei-bianjie/msg-parser/types"
	aTypes "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"strings"
	"time"
)

var (
	_parser msgparser.MsgParser
	_conf   *config.Config
)

func InitRouter(conf *config.Config) {
	_conf = conf
	initBech32Prefix(conf)
	router := msgparser.RegisteRouter()
	if conf.Server.OnlySupportModule != "" {
		modules := strings.Split(conf.Server.OnlySupportModule, ",")
		msgRoute := msgparser.NewRouter()
		for _, one := range modules {
			fn, exist := msgparser.RouteHandlerMap[one]
			if !exist {
				logger.Fatal("no support module: " + one)
			}
			msgRoute = msgRoute.AddRoute(one, fn)
		}
		if msgRoute.GetRoutesLen() > 0 {
			router = msgRoute
		}

	}
	_parser = msgparser.NewMsgParser(router)
	_conf = conf
}

func ParseBlockAndTxs(b int64, client *pool.Client) (*models.Block, []*models.Tx, []txn.Op, error) {
	var (
		blockDoc models.Block
		block    *ctypes.ResultBlock
		txnOps   []txn.Op
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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

	txResultMap := handleTxResult(client, block.Block)

	txDocs := make([]*models.Tx, 0, len(block.Block.Txs))
	if len(block.Block.Txs) > 0 {
		for _, v := range block.Block.Txs {
			txHash := utils.BuildHex(v.Hash())
			txResult, ok := txResultMap[txHash]
			if !ok || txResult.TxResult == nil {
				return &blockDoc, txDocs, txnOps, utils.ConvertErr(block.Block.Height, txHash, "TxResult",
					fmt.Errorf("no found"))
			}
			if txResult.Err != nil {
				return &blockDoc, txDocs, txnOps, utils.ConvertErr(block.Block.Height, txHash, "TxResult",
					txResult.Err)
			}
			txDoc, ops, err := parseTx(v, txResult.TxResult, block.Block)
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

func parseTx(txBytes types.Tx, txResult *ctypes.ResultTx, block *types.Block) (models.Tx, []txn.Op, error) {
	var (
		docTx models.Tx

		docTxMsgs []msgsdktypes.TxMsg
		txnOps    []txn.Op
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

		if msgDocInfo.DocTxMsg.Type == MsgTypeEthereumTx {
			var msgEtheumTx evm.DocMsgEthereumTx
			var txData msgparser.LegacyTx
			utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(msgDocInfo.DocTxMsg.Msg), &msgEtheumTx)
			utils.UnMarshalJsonIgnoreErr(msgEtheumTx.Data, &txData)
			docTx.ContractAddrs = append(docTx.ContractAddrs, txData.To)
			docTx.Fee.Gas = min(txResult.TxResult.GasUsed, docTx.Fee.Gas)
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
	docTx.ContractAddrs = removeDuplicatesFromSlice(docTx.ContractAddrs)
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
func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
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

func SaveDocsWithTxn(blockDoc *models.Block, txDocs []*models.Tx, taskDoc models.SyncTask, opsDoc []txn.Op) error {
	var (
		insertOps []txn.Op
	)

	if blockDoc.Height == 0 {
		return fmt.Errorf("invalid block, height equal 0")
	}

	blockOp := txn.Op{
		C:      models.BlockModel.Name(),
		Id:     bson.NewObjectId(),
		Insert: blockDoc,
	}

	dataLen := 0
	if length := len(txDocs); length > 0 {

		insertOps = make([]txn.Op, 0, _conf.Server.InsertBatchLimit)
		for _, v := range txDocs {
			op := txn.Op{
				C:      models.TxModel.Name(),
				Id:     bson.NewObjectId(),
				Insert: v,
			}
			dataLen += 1
			if dataLen >= _conf.Server.InsertBatchLimit {
				if err := models.Txn(insertOps); err != nil {
					return err
				}
				insertOps = make([]txn.Op, 0, _conf.Server.InsertBatchLimit)
				insertOps = append(insertOps, op)
				dataLen = 0
			} else {
				insertOps = append(insertOps, op)
			}
		}
	}
	if taskDoc.ID.Valid() {
		updateOp := txn.Op{
			C:      models.SyncTaskModel.Name(),
			Id:     taskDoc.ID,
			Assert: txn.DocExists,
			Update: bson.M{
				"$set": bson.M{
					"current_height":   taskDoc.CurrentHeight,
					"status":           taskDoc.Status,
					"last_update_time": taskDoc.LastUpdateTime,
				},
			},
		}
		insertOps = append(insertOps, updateOp)
	}

	//ops = append(append(ops, blockOp), binanceTxsOps...)
	insertOps = append(insertOps, blockOp)
	if len(opsDoc) > 0 {
		insertOps = append(insertOps, opsDoc...)
	}

	if len(insertOps) > 0 {
		err := models.Txn(insertOps)
		if err != nil {
			return err
		}
	}

	return nil
}
