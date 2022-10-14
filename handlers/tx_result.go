package handlers

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	ctypes "github.com/okex/exchain/libs/tendermint/rpc/core/types"
	"github.com/okex/exchain/libs/tendermint/types"
	"time"
)

type chanTxResult struct {
	TxHash   string
	TxResult *ctypes.ResultTx
	Err      error
}

// parse tx with more goroutine concurrency
func handleTxResult(client *pool.Client, block *types.Block) map[string]chanTxResult {
	if _conf == nil {
		logger.Fatal("InitRouter don't work")
	}
	chanParseTxLimit := make(chan bool, 5)
	chanRes := make(chan chanTxResult, len(block.Txs))
	for _, v := range block.Txs {
		chanParseTxLimit <- true
		// parse txReult with more goroutine concurrency
		go getTxResult(client, v, chanParseTxLimit, chanRes)
	}
	txRetMap := make(map[string]chanTxResult, len(block.Txs))
	for i := 0; i < len(block.Txs); i++ {
		chanValue := <-chanRes
		txRetMap[chanValue.TxHash] = chanValue
	}
	return txRetMap
}
func includeIbcTxs(txBytes types.Tx) bool {
	var inclueIbcTx bool
	authTx, err := codec.GetStdTx(txBytes)
	if err != nil {
		logger.Warn(err.Error())
		return inclueIbcTx
	}
	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return inclueIbcTx
	}
	for _, v := range msgs {
		msgDocInfo := _parser.HandleTxMsg(v)
		_, ok := _filterMap[msgDocInfo.DocTxMsg.Type]
		if ok {
			return true
		}
	}
	return inclueIbcTx
}

func getTxResult(c *pool.Client, txBytes types.Tx, chanLimit chan bool, chanRes chan chanTxResult) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("execute getTxResult fail", logger.Any("err", r))
		}
		<-chanLimit
	}()
	var (
		txResult *ctypes.ResultTx
		err      error
	)
	if includeIbcTxs(txBytes) {
		txResult, err = c.Tx(txBytes, false)
		if err != nil {
			time.Sleep(1 * time.Second)
			if v, err1 := c.Tx(txBytes, false); err1 != nil {
				err = err1
			} else {
				txResult = v
			}
		}
	}

	if txResult == nil {
		chanRes <- chanTxResult{Err: err}
		return
	}
	ret := chanTxResult{
		TxHash:   txResult.Hash.String(),
		TxResult: txResult,
		Err:      err,
	}
	chanRes <- ret

	return
}
