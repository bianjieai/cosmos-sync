package handlers

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	aTypes "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"gitlab.bianjie.ai/irita/ex-sync/libs/cdc"
	"gitlab.bianjie.ai/irita/ex-sync/libs/logger"
	"gitlab.bianjie.ai/irita/ex-sync/libs/pool"
	"gitlab.bianjie.ai/irita/ex-sync/models"
	mMsg "gitlab.bianjie.ai/irita/ex-sync/models/msgs"
	itypes "gitlab.bianjie.ai/irita/ex-sync/types"
	"gitlab.bianjie.ai/irita/ex-sync/utils"
	"gitlab.bianjie.ai/irita/ex-sync/utils/constant"
	"time"
)

func ParseBlockAndTxs(b int64, client *pool.Client) (*models.Block, []*models.Tx, error) {
	var (
		blockDoc models.Block
		block    *ctypes.ResultBlock
	)

	if v, err := client.Block(&b); err != nil {
		logger.Warn("parse block fail, now try again", logger.Int64("height", b),
			logger.String("err", err.Error()))
		if v2, err := client.Block(&b); err != nil {
			logger.Error("parse block fail", logger.Int64("height", b),
				logger.String("err", err.Error()))
			return &blockDoc, nil, err
		} else {
			block = v2
		}
	} else {
		block = v
	}
	blockDoc = models.Block{
		Height: block.Block.Height,
		Time:   block.Block.Time,
		Hash:   block.BlockMeta.Header.Hash().String(),
		Txn:    block.BlockMeta.Header.NumTxs,
	}

	txDocs := make([]*models.Tx, 0, len(block.Block.Txs))
	if len(block.Block.Txs) > 0 {
		for _, v := range block.Block.Txs {
			txDoc := parseTx(client, v, block.Block.Time)
			if txDoc.TxHash != "" {
				txDocs = append(txDocs, &txDoc)
			}
		}
	}

	return &blockDoc, txDocs, nil
}

func parseTx(c *pool.Client, txBytes types.Tx, blockTime time.Time) models.Tx {
	var (
		stdTx auth.StdTx
		docTx models.Tx

		complexMsg       bool
		txType, from, to string
		coins            []models.Coin
		signer           string
		docTxMsgs        []models.DocTxMsg
		signers          []string
	)

	if txResult, err := c.Tx(txBytes.Hash(), false); err != nil {
		logger.Error("get tx result fail", logger.String("txHash", txBytes.String()),
			logger.String("err", err.Error()))
		return docTx
	} else {
		docTx.Time = blockTime
		docTx.Height = txResult.Height
		docTx.TxHash = utils.BuildHex(txBytes.Hash())
		docTx.Status = parseTxStatus(txResult.TxResult.Code)
		docTx.Log = txResult.TxResult.Log
		docTx.Events = parseEvents(txResult.TxResult.Events)

		if err := cdc.Cdc.UnmarshalBinaryLengthPrefixed(txResult.Tx, &stdTx); err != nil {
			logger.Error("unmarshal tx fail", logger.String("txHash", docTx.TxHash),
				logger.String("err", err.Error()))
			return docTx
		}

		docTx.Memo = stdTx.Memo

		msgs := stdTx.GetMsgs()
		if len(msgs) == 0 {
			return docTx
		}
		for i, v := range msgs {
			var (
				msgDocInfo mMsg.MsgDocInfo
			)
			switch v.(type) {
			case itypes.MsgSend:
				docMsg := mMsg.DocMsgSend{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgSend))
				break
			case itypes.MsgNFTMint:
				docMsg := mMsg.DocMsgNFTMint{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgNFTMint))
				break
			case itypes.MsgNFTEdit:
				docMsg := mMsg.DocMsgNFTEdit{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgNFTEdit))
				break
			case itypes.MsgNFTTransfer:
				docMsg := mMsg.DocMsgNFTTransfer{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgNFTTransfer))
				break
			case itypes.MsgNFTBurn:
				docMsg := mMsg.DocMsgNFTBurn{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgNFTBurn))
				break
			case itypes.MsgServiceDef:
				docMsg := mMsg.DocMsgServiceDef{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgServiceDef))
				break
			case itypes.MsgServiceBind:
				docMsg := mMsg.DocMsgServiceBind{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgServiceBind))
				break
			case itypes.MsgServiceRequest:
				docMsg := mMsg.DocMsgServiceRequest{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgServiceRequest))
				break
			case itypes.MsgServiceResponse:
				docMsg := mMsg.DocMsgServiceResponse{}
				msgDocInfo = docMsg.HandleTxMsg(v.(itypes.MsgServiceResponse))
				break
			}

			if msgDocInfo.Signer == "" {
				continue
			}

			if !complexMsg {
				complexMsg = msgDocInfo.ComplexMsg
			}
			if i == 0 {
				txType = msgDocInfo.DocTxMsg.Type
				from = msgDocInfo.From
				to = msgDocInfo.To
				coins = msgDocInfo.Coins
				signer = msgDocInfo.Signer
			}
			docTxMsgs = append(docTxMsgs, msgDocInfo.DocTxMsg)
			signers = msgDocInfo.Signers
		}

		if !complexMsg && len(msgs) > 1 {
			complexMsg = true
		}

		docTx.ComplexMsg = complexMsg
		docTx.Type = txType
		docTx.From = from
		docTx.To = to
		docTx.Coins = coins
		docTx.Signer = signer
		docTx.DocTxMsgs = docTxMsgs
		docTx.Signers = signers

		// don't save txs which have not parsed
		if docTx.Type == "" {
			return models.Tx{}
		}

		return docTx
	}
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
