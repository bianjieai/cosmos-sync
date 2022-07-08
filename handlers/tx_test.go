package handlers

import (
	"context"
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/tendermint/tendermint/types"
	"testing"
	"time"
)

func TestParseTxs(t *testing.T) {

	block := int64(15620522)
	//block := int64(3705556) //sentinel
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

	if blockDoc, txDocs, err := ParseBlockAndTxs(block, c); err != nil {
		t.Fatal(err)
	} else {
		t.Log(utils.MarshalJsonIgnoreErr(blockDoc))
		t.Log(utils.MarshalJsonIgnoreErr(txDocs))
	}
}

func Test_handleTxResult(t *testing.T) {
	var block *types.Block
	b := int64(9179568)
	//block := int64(3705556) //sentinel
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

	ctx := context.Background()
	if v, err := c.Block(ctx, &b); err != nil {
		time.Sleep(1 * time.Second)
		if v2, err := c.Block(ctx, &b); err != nil {
			t.Fatal(err.Error())
		} else {
			block = v2.Block
		}
	} else {
		block = v.Block
	}

	mapData := handleTxResult(c, block)
	t.Log(utils.MarshalJsonIgnoreErr(mapData))

}

func TestUnmarshalTibcAckInEvents(t *testing.T) {
	bytesdata := []byte("\ufffd\u0001\u0001\u0001")
	var result TIBCAcknowledgement
	err := result.Unmarshal(bytesdata)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(result.String())
}
