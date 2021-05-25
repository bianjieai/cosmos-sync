package pool

import (
	"context"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	jsonrpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
)

type BlockChainClient struct {
	remote string
	client *jsonrpcclient.Client
}

func (c *BlockChainClient) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	result := new(ctypes.ResultBlock)
	params := make(map[string]interface{})
	if height != nil {
		params["height"] = height
	}
	_, err := c.client.Call(ctx, "block", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *BlockChainClient) Status(ctx context.Context) (*ctypes.ResultStatus, error) {
	result := new(ctypes.ResultStatus)
	_, err := c.client.Call(ctx, "status", map[string]interface{}{}, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *BlockChainClient) Tx(ctx context.Context, hash []byte, prove bool) (*ctypes.ResultTx, error) {
	result := new(ctypes.ResultTx)
	params := map[string]interface{}{
		"hash":  hash,
		"prove": prove,
	}
	_, err := c.client.Call(ctx, "tx", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
