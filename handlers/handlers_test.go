package handlers

import (
	"context"
	"github.com/bianjieai/irita-sync/libs/pool"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/kaifei-bianjie/msg-parser/codec"
	"testing"
)

func TestHandleTxMsg(t *testing.T) {

	block := int64(106329)
	client := pool.GetClient()
	defer func() {
		client.Release()
	}()
	ctx := context.Background()
	v, err := client.Block(ctx, &block)
	if err != nil {

	}
	for _, txBytes := range v.Block.Txs {
		authTx, err := codec.GetSigningTx(txBytes)
		if err != nil {

		}
		msgs := authTx.GetMsgs()
		for _, msg := range msgs {
			msgDocInfo, ops := HandleTxMsg(msg)
			t.Log(utils.MarshalJsonIgnoreErr(msgDocInfo))
			t.Log(utils.MarshalJsonIgnoreErr(ops))
		}
	}
}
