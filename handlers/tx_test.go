package handlers

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/utils"
	"testing"
)

func TestParseTxs(t *testing.T) {
	block := int64(1097850)
	conf, err := config.ReadConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	InitRouter(conf)
	pool.Init(conf)
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()

	if blockDoc, txDocs, _, err := ParseBlockAndTxs(block, c); err != nil {
		t.Fatal(err)
	} else {
		t.Log(utils.MarshalJsonIgnoreErr(blockDoc))
		t.Log(utils.MarshalJsonIgnoreErr(txDocs))

		//b, _ := hex.DecodeString("736572766963652063616c6c20726573706f6e7365")
		//t.Log(string(b))
	}
}
