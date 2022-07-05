package pool

import (
	"context"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (c *Client) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	if isJsonRpcProtocol {
		return c.Jrpc.Block(ctx, height)
	}
	http := c.HTTP
	return http.Block(ctx, height)

}

func (c *Client) Tx(ctx context.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	if isJsonRpcProtocol {
		return c.Jrpc.Tx(ctx, hash, prove)
	}
	http := c.HTTP
	return http.Tx(ctx, hash, prove)

}

func (c *Client) BlockResults(ctx context.Context, height *int64) (*ctypes.ResultBlockResults, error) {
	if isJsonRpcProtocol {
		return c.Jrpc.BlockResults(ctx, height)
	}
	http := c.HTTP
	return http.BlockResults(ctx, height)
}
