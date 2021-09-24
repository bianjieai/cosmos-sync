package pool

import (
	"context"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (c *Client) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	if useJrpc {
		return c.Jrpc.Block(ctx, height)
	}
	http := c.HTTP
	return http.Block(ctx, height)

}

func (c *Client) Tx(ctx context.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	if useJrpc {
		return c.Jrpc.Tx(ctx, hash, prove)
	}
	http := c.HTTP
	return http.Tx(ctx, hash, prove)

}
