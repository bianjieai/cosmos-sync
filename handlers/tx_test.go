package handlers

import (
	"gitlab.bianjie.ai/irita/ex-sync/libs/pool"
	"gitlab.bianjie.ai/irita/ex-sync/utils"
	"testing"
)

func TestParseTxs(t *testing.T) {
	block := int64(52344944)
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
