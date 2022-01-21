package handlers

import (
	"context"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/utils"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"sync"
	"time"
)

// parse tx with more goroutine concurrency
func handleTxResult(client *pool.Client, block *types.Block) map[string]*ctypes.ResultTx {
	if _conf == nil {
		logger.Fatal("InitRouter don't work")
	}
	txRetMap := make(map[string]*ctypes.ResultTx, len(block.Txs))
	if _conf.Server.ThreadNumParseTx <= 0 {
		_conf.Server.ThreadNumParseTx = 1
	}

	mutex := &sync.Mutex{}
	group := &sync.WaitGroup{}
	chanParseTxLimit := make(chan bool, _conf.Server.ThreadNumParseTx)
	for _, v := range block.Txs {
		chanParseTxLimit <- true
		group.Add(1)
		// parse txReult with more goroutine concurrency
		go getTxResult(client, v, block.Height, chanParseTxLimit, txRetMap, mutex, group)
	}
	group.Wait()
	return txRetMap
}

func getTxResult(c *pool.Client, txBytes types.Tx, height int64, chanLimit chan bool, txRetMap map[string]*ctypes.ResultTx,
	mutex *sync.Mutex, group *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("execute getTxResult fail", logger.Any("err", r))
		}
		group.Done()
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
	logger.Debug("get txResult ok", logger.String("txHash", txHash))
	mutex.Lock()
	txRetMap[txHash] = txResult
	mutex.Unlock()
	return
}
