package handlers

import (
	"context"
	"fmt"
	"github.com/bianjieai/irita-sync/libs/cdc"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/libs/pool"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/bianjieai/irita-sync/utils/constant"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	aTypes "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"gopkg.in/mgo.v2/txn"
	"strings"
	"time"
)

func ParseBlockAndTxs(b int64, client *pool.Client) (*models.Block, []*models.Tx, []txn.Op, error) {
	var (
		blockDoc models.Block
		block    *ctypes.ResultBlock
		txnOps   []txn.Op
	)

	ctx := context.Background()
	if v, err := client.Block(ctx, &b); err != nil {
		time.Sleep(500 * time.Millisecond)
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
		docTx     models.Tx
		docTxMsgs []models.DocTxMsg
		txnOps    []txn.Op
		log       string
	)

	txHash := utils.BuildHex(txBytes.Hash())
	height := block.Height
	Tx, err := cdc.GetTxDecoder()(txBytes)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("txhash", txHash),
			logger.Int64("height", height))
		return docTx, txnOps, nil
	}

	authTx := Tx.(signing.Tx)
	fee := models.BuildFee(authTx.GetFee(), authTx.GetGas())
	memo := authTx.GetMemo()
	ctx := context.Background()
	txResult, err := c.Tx(ctx, txBytes.Hash(), false)
	if err != nil {
		time.Sleep(500 * time.Millisecond)
		if ret, err := c.Tx(ctx, txBytes.Hash(), false); err != nil {
			return docTx, txnOps, utils.ConvertErr(height, txHash, "TxResult", err)
		} else {
			txResult = ret
		}
	}
	status := parseTxStatus(txResult.TxResult.Code)
	if status == constant.TxStatusFail {
		log = txResult.TxResult.Log
	}
	txIndex := txResult.Index
	docTx = models.Tx{
		Height:  height,
		Time:    block.Time.Unix(),
		TxHash:  txHash,
		Fee:     &fee,
		Memo:    memo,
		Status:  status,
		Log:     log,
		Events:  parseEvents(txResult.TxResult.Events),
		TxIndex: txIndex,
	}
	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return docTx, txnOps, nil
	}

	for i, v := range msgs {
		msgDocInfo, ops := HandleTxMsg(v)
		if len(msgDocInfo.Addrs) == 0 {
			continue
		}
		if i == 0 {
			docTx.Type = msgDocInfo.DocTxMsg.Type
		}
		for _, signer := range v.GetSigners() {
			docTx.Signers = append(docTx.Signers, signer.String())
		}

		docTx.Addrs = append(docTx.Addrs, removeDuplicatesFromSlice(msgDocInfo.Addrs)...)
		docTxMsgs = append(docTxMsgs, msgDocInfo.DocTxMsg)
		docTx.Types = append(docTx.Types, msgDocInfo.DocTxMsg.Type)
		if len(ops) > 0 {
			txnOps = append(txnOps, ops...)
		}
	}
	docTx.Signers = removeDuplicatesFromSlice(docTx.Signers)
	docTx.Types = removeDuplicatesFromSlice(docTx.Types)
	docTx.Addrs = removeDuplicatesFromSlice(docTx.Addrs)

	docTx.DocTxMsgs = docTxMsgs

	// don't save txs which have not parsed
	if docTx.Type == "" {
		logger.Warn(constant.NoSupportMsgTypeTag,
			logger.String("txhash", txHash),
			logger.Int64("height", height))
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
