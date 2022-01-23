package handlers

import (
	"context"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/utils"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"time"
)

type chanTxResult struct {
	TxHash   string
	TxResult *ctypes.ResultTx
}

// parse tx with more goroutine concurrency
func handleTxResult(client *pool.Client, block *types.Block) map[string]*ctypes.ResultTx {
	if _conf == nil {
		logger.Fatal("InitRouter don't work")
	}
	if _conf.Server.ThreadNumParseTx <= 0 {
		_conf.Server.ThreadNumParseTx = 1
	}

	chanParseTxLimit := make(chan bool, _conf.Server.ThreadNumParseTx)
	chanRes := make(chan chanTxResult, len(block.Txs))
	for _, v := range block.Txs {
		chanParseTxLimit <- true
		// parse txReult with more goroutine concurrency
		go getTxResult(client, v, block.Height, chanParseTxLimit, chanRes)
	}
	txRetMap := make(map[string]*ctypes.ResultTx, len(block.Txs))
	for i := 0; i < len(block.Txs); i++ {
		chanValue := <-chanRes
		txRetMap[chanValue.TxHash] = chanValue.TxResult

	}
	return txRetMap
}

func getTxResult(c *pool.Client, txBytes types.Tx, height int64, chanLimit chan bool, chanRes chan chanTxResult) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("execute getTxResult fail", logger.Any("err", r))
		}
		<-chanLimit
	}()
	txHash := utils.BuildHex(txBytes.Hash())
	ctx := context.Background()
	txResult, err := c.Tx(ctx, txBytes.Hash(), false)
	if err != nil {
		time.Sleep(1 * time.Second)
		if v, err := c.Tx(ctx, txBytes.Hash(), false); err != nil {
			logger.Error(utils.ConvertErr(height, txHash, "TxResult", err).Error())
		} else {
			txResult = v
		}
	}

	ret := chanTxResult{
		TxHash:   txResult.Hash.String(),
		TxResult: txResult,
	}
	chanRes <- ret

	return
}
