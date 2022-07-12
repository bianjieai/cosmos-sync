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
	"github.com/kaifei-bianjie/msg-parser/modules/mt"
	msgsdktypes "github.com/kaifei-bianjie/msg-parser/types"
	types2 "github.com/tendermint/tendermint/abci/types"
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

	blockResults, err := client.BlockResults(context.Background(), &b)
	if err != nil {
		time.Sleep(1 * time.Second)
		blockResults, err = client.BlockResults(context.Background(), &b)
		if err != nil {
			return &blockDoc, nil, txnOps, utils.ConvertErr(b, "", "ParseBlockResult", err)
		}
	}

	if len(block.Block.Txs) != len(blockResults.TxsResults) {
		return nil, nil, nil, utils.ConvertErr(b, "", "block.Txs length not equal blockResult", nil)
	}

	txDocs := make([]*models.Tx, 0, len(block.Block.Txs))
	if len(block.Block.Txs) > 0 {
		for i, v := range block.Block.Txs {
			txResult := blockResults.TxsResults[i]
			txDoc, ops, err := parseTx(uint32(i), v, txResult, block.Block)
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

func parseTx(index uint32, txBytes types.Tx, txResult *types2.ResponseDeliverTx, block *types.Block) (models.Tx, []txn.Op, error) {
	var (
		docTx models.Tx

		docTxMsgs []msgsdktypes.TxMsg
		txnOps    []txn.Op
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
	docTx.TxIndex = index

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("errTag", "TxDecoder"),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return docTx, txnOps, nil
	}
	docTx.GasUsed = txResult.GasUsed
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
		}

		if msgDocInfo.DocTxMsg.Type == MsgTypeMTIssueDenom && docTx.Status == constant.TxStatusSuccess {
			// get denom_id from events then set to msg, because this msg hasn't denom_id
			denomId := ParseAttrValueFromEvents(docTx.EventsNew[i].Events, EventTypeIssueDenom, AttrKeyDenomId)
			msg := msgDocInfo.DocTxMsg.Msg.(*mt.DocMsgMTIssueDenom)
			msg.Id = denomId
			msgDocInfo.DocTxMsg.Msg = msg
		}

		if msgDocInfo.DocTxMsg.Type == MsgTypeMintMT && docTx.Status == constant.TxStatusSuccess {
			// get mt_id from events then set to msg, because this msg hasn't mt_id
			mtId := ParseAttrValueFromEvents(docTx.EventsNew[i].Events, EventTypeMintMT, AttrKeyMTId)
			msg := msgDocInfo.DocTxMsg.Msg.(*mt.DocMsgMTMint)
			msg.Id = mtId
			msgDocInfo.DocTxMsg.Msg = msg
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
func parseTxStatus(code uint32) uint32 {
	if code == 0 {
		return constant.TxStatusSuccess
	} else {
		return constant.TxStatusFail
	}
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
