package handlers

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	"testing"
)

func Test_handleTxResult(t *testing.T) {
	conf, err := config.ReadConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	models.Init(conf)
	InitRouter(conf)
	pool.Init(conf)
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()
	b := int64(14121638)
	block, err := c.Block(&b)
	if err != nil {
		t.Fatal(err.Error())
	}
	res := handleTxResult(c, block.Block)
	for _, val := range res {
		t.Log(val.TxHash, utils.MarshalJsonIgnoreErr(val.TxResult))
	}
}
