package handlers

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/msgparser"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/evm"
	msgtypes "github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	types2 "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	"time"
)

var _parser msgparser.MsgParser

func InitMsgParser() {
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
		logger.Warn(err.Error(),
			logger.String("errTag", "TxDecoder"),
			logger.String("txhash", txHash),
			logger.Int64("height", height))
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
		GasUsed: txResult.GasUsed,
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
		case MsgTypeEthereumTx:
			var msgEtheumTx evm.DocMsgEthereumTx
			var txData msgparser.LegacyTx
			utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(msgDocInfo.DocTxMsg.Msg), &msgEtheumTx)
			utils.UnMarshalJsonIgnoreErr(msgEtheumTx.Data, &txData)
			docTx.ContractAddrs = append(docTx.ContractAddrs, txData.To)
			if len(txResult.Data) > 0 {
				if txRespond, err := evmtypes.DecodeTxResponse(txResult.Data); err == nil {
					if len(txRespond.Ret) > 0 {
						docTx.EvmTxRespondRet = hex.EncodeToString(txRespond.Ret)
					}
				} else {
					logger.Warn("DecodeTxResponse failed",
						logger.String("err", err.Error()),
						logger.String("txhash", txHash),
						logger.Int64("height", block.Height))
				}
			}
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

	feeGranter := authTx.FeeGranter()
	if feeGranter != nil {
		docTx.FeeGranter = feeGranter.String()
	}
	feePayer := authTx.FeePayer()
	if feePayer != nil {
		docTx.FeePayer = feePayer.String()
	}
	feeGrantee := GetFeeGranteeFromEvents(txResult.Events)
	if feeGrantee != "" {
		docTx.FeeGrantee = feeGrantee
	}

	// don't save txs which have not parsed
	if docTx.Type == "" {
		logger.Warn(constant.NoSupportMsgTypeTag,
			logger.String("errTag", "TxMsg"),
			logger.String("txhash", txHash),
			logger.Int64("height", height))
		return models.Tx{}, nil
	}

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

func GetFeeGranteeFromEvents(events []types2.Event) string {
	for _, val := range events {
		if val.Type == constant.UseGrantee || val.Type == constant.SetGrantee {
			for _, attribute := range val.Attributes {
				if fmt.Sprintf("%s", attribute.Key) == constant.Grantee {
					return fmt.Sprintf("%s", attribute.Value)
				}
			}
		}
	}
	return ""
}
