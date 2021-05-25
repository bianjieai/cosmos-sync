package pool

import (
	"context"
	"encoding/hex"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"testing"
)

var c *Client
var err error

func TestMain(m *testing.M) {
	c, err = newClient("http://localhost:26657")
	if err != nil {
		panic(err)
	}

	m.Run()
}

func TestBlock(t *testing.T) {
	var height int64 = 1
	block, err := c.Block(context.Background(), &height)
	if err != nil {
		panic(err)
	}
	t.Log(block)
}

func TestStatus(t *testing.T) {
	status, err := c.Status(context.Background())
	if err != nil {
		panic(err)
	}
	t.Log(status)
}

func TestTx(t *testing.T) {
	txHashBytes, err := hex.DecodeString("79560EBFD400F438079044DA28F99C9F5710089BA12E46975C16538471F6E18E")
	if err != nil {
		panic(err)
	}

	resultTx, err := c.Tx(context.Background(), txHashBytes, false)
	if err != nil {
		panic(err)
	}
	t.Log(resultTx)
}

func TestHeart(t *testing.T) {
	result := new(ctypes.ResultStatus)
	resultHeart, err := c.client.Call(ctx, "health", map[string]interface{}{}, result)
	if err != nil {
		panic(err)
	}
	t.Log(resultHeart)
}
