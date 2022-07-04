package jrpc

import (
	"context"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"strconv"
)

type BlockChainClient struct {
	remote string
	client *JsonRpcClient
}

func (c *BlockChainClient) Block(ctx context.Context, height *int64) (*ctypes.ResultBlock, error) {
	result := new(ctypes.ResultBlock)
	params := make(map[string]interface{})
	if height != nil {
		params["height"] = strconv.FormatInt(*height, 10)
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

func (c *BlockChainClient) BlockResults(
	ctx context.Context,
	height *int64,
) (*ctypes.ResultBlockResults, error) {
	result := new(ctypes.ResultBlockResults)
	params := make(map[string]interface{})
	if height != nil {
		params["height"] = strconv.FormatInt(*height, 10)
	}
	_, err := c.client.Call(ctx, "block_results", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *BlockChainClient) HeartBeat(ctx context.Context) error {
	result := new(ctypes.ResultStatus)
	_, err := c.client.Call(ctx, "health", map[string]interface{}{}, result)
	return err
}

func NewBlockChainClient(nodeUrl string) (*BlockChainClient, error) {
	customClient, err := NewJsonRpcClient(nodeUrl)

	blockChainClient := &BlockChainClient{
		remote: nodeUrl,
		client: customClient,
	}
	return blockChainClient, err
}
