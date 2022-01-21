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

// parse tx with more goroutine concurrency
func handleTxResult(client *pool.Client, block *types.Block) map[string]ctypes.ResultTx {
	if _conf == nil {
		logger.Fatal("InitRouter don't work")
	}
	txRetMap := make(map[string]ctypes.ResultTx, len(block.Txs))
	if _conf.Server.ThreadNumParseTx <= 0 {
		_conf.Server.ThreadNumParseTx = 1
	}

	chanParseTxLimit := make(chan bool, _conf.Server.ThreadNumParseTx)
	for _, v := range block.Txs {
		chanParseTxLimit <- true
		var txResult ctypes.ResultTx
		// parse txReult with more goroutine concurrency
		go getTxResult(client, v, block.Height, chanParseTxLimit, &txResult)
		txHash := utils.BuildHex(v.Hash())
		logger.Debug("get txResult ok", logger.String("txHash", txHash))
		txRetMap[txHash] = txResult
	}

	return txRetMap
}

func getTxResult(c *pool.Client, txBytes types.Tx, height int64, chanLimit chan bool, ret *ctypes.ResultTx) {
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
			ret = v
		}
	} else {
		ret = txResult
	}

	return
}
