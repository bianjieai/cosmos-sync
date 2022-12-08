package handlers

import (
	"context"
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/msgparser"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	common_parser "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/common-parser/codec"
	msgtypes "github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	"github.com/kaifei-bianjie/irismod-parser/modules/mt"
	types2 "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"strings"
	"time"
)

var _parser msgparser.MsgParser

func InitRouter(conf *config.Config) {
	if conf.Server.OnlySupportModule != "" {
		resRouteClient := make(map[string]common_parser.Client, 0)
		modules := strings.Split(conf.Server.OnlySupportModule, ",")
		for _, one := range modules {
			fn, exist := msgparser.RouteClientMap[one]
			if !exist {
				logger.Fatal("no support module: " + one)
			}
			resRouteClient[one] = fn
		}
		if len(resRouteClient) > 0 {
			msgparser.RouteClientMap = resRouteClient
		}
	}
	_parser = msgparser.NewMsgParser()
}

func ParseBlockAndTxs(b int64, client *pool.Client) (*models.Block, []*models.Tx, error) {
	var (
		blockDoc models.Block
		block    *ctypes.ResultBlock
	)
	ctx := context.Background()
	if v, err := client.Block(ctx, &b); err != nil {
		time.Sleep(500 * time.Millisecond)
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

	if blockDoc.Txn <= 0 {
		return &blockDoc, nil, nil
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
		docTx     models.Tx
		docTxMsgs []msgtypes.TxMsg
		log       string
	)

	txHash := utils.BuildHex(txBytes.Hash())
	height := block.Height
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		if docTx.Status == constant.TxStatusFail {
			docTx.Type = constant.NoSupportModule
			docTx.DocTxMsgs = append(docTx.DocTxMsgs, msgtypes.TxMsg{
				Type: docTx.Type,
			})
			return docTx, nil
		}
		for i := range docTx.EventsNew {
			msgName := ParseAttrValueFromEvents(docTx.EventsNew[i].Events, EventTypeMessage, AttrKeyAction)
			module := _parser.GetModule(msgName)
			_, exist := msgparser.RouteClientMap[module]
			if docTx.EventsNew[i].MsgIndex == 0 {
				if !exist {
					docTx.Type = constant.NoSupportModule
				} else {
					docTx.Type = constant.IncorrectParse
				}
				docTx.DocTxMsgs = append(docTx.DocTxMsgs, msgtypes.TxMsg{
					Type: docTx.Type,
				})
			}
			docTx.Types = append(docTx.Types, msgName)
		}
		docTx.Types = removeDuplicatesFromSlice(docTx.Types)

		//logger.Warn(err.Error(),
		//	logger.String("errTag", "TxDecoder"),
		//	logger.String("txhash", txHash),
		//	logger.Int64("height", block.Height))
		return docTx, nil
	}
	fee := msgtypes.BuildFee(authTx.GetFee(), authTx.GetGas())
	memo := authTx.GetMemo()
	status := parseTxStatus(txResult.Code)
	if status == constant.TxStatusFail {
		log = txResult.Log
	}
	docTx = models.Tx{
		Height:  height,
		Time:    block.Time.Unix(),
		TxHash:  txHash,
		Fee:     fee,
		Memo:    memo,
		Status:  status,
		Log:     log,
		TxIndex: uint32(index),
		TxId:    block.Height*100000 + int64(index),
	}
	docTx.EventsNew = parseABCILogs(txResult.Log)
	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return docTx, nil
	}

	for i, v := range msgs {
		msgDocInfo := _parser.HandleTxMsg(v)
		if len(msgDocInfo.Addrs) == 0 {
			continue
		}
		if i == 0 {
			docTx.Type = msgDocInfo.DocTxMsg.Type
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
		}

		for _, signer := range v.GetSigners() {
			docTx.Signers = append(docTx.Signers, signer.String())
		}
		docTx.Addrs = append(docTx.Addrs, removeDuplicatesFromSlice(msgDocInfo.Addrs)...)
		docTxMsgs = append(docTxMsgs, msgDocInfo.DocTxMsg)
		docTx.Types = append(docTx.Types, msgDocInfo.DocTxMsg.Type)
	}
	docTx.Signers = removeDuplicatesFromSlice(docTx.Signers)
	docTx.Types = removeDuplicatesFromSlice(docTx.Types)
	docTx.Addrs = removeDuplicatesFromSlice(docTx.Addrs)

	docTx.DocTxMsgs = docTxMsgs

	//// don't save txs which have not parsed
	//if docTx.Type == "" {
	//	logger.Warn(constant.NoSupportMsgTypeTag,
	//		logger.String("errTag", "TxMsg"),
	//		logger.String("txhash", txHash),
	//		logger.Int64("height", height))
	//	return models.Tx{}, nil
	//}

	return docTx, nil
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

const (
	EventTypeIssueDenom = "issue_denom"
	EventTypeMintMT     = "mint_mt"
	AttrKeyDenomId      = "denom_id"
	AttrKeyMTId         = "mt_id"
	EventTypeMessage    = "message"
	AttrKeyAction       = "action"
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
