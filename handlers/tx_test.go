package handlers

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"testing"
)

func TestParseTxs(t *testing.T) {
	block := int64(3399387)
	conf, err := config.ReadConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	//models.Init(conf)
	models.SetConfig(conf)
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

func TestUnmarshalTibcAckInEvents(t *testing.T) {
	bytesdata := []byte("\ufffd\u0001\u0001\u0001")
	var result TIBCAcknowledgement
	err := result.Unmarshal(bytesdata)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(result.String())
}
