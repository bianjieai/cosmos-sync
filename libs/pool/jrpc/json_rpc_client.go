package jrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"io/ioutil"
	"net/http"
)

type JsonRpcClient struct {
	address  string
	username string
	password string
}

func NewJsonRpcClient(nodeUrl string) (*JsonRpcClient, error) {
	if len(nodeUrl) == 0 {
		return nil, fmt.Errorf("nodeUrl is empty")
	}
	return &JsonRpcClient{
		address: nodeUrl,
	}, nil
}

func (c *JsonRpcClient) Call(ctx context.Context, method string,
	params map[string]interface{}, result interface{}) (interface{}, error) {
	requestBytes, err := c.mapToRequest(method, params)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	httpResponse, err := http.Post(c.address, "application/json", bytes.NewReader(requestBytes))
	if err != nil {
		return nil, fmt.Errorf("post failed: %w", err)
	}
	defer httpResponse.Body.Close()

	httpResponseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	rpcResponse := &types.RPCResponse{}
	if err = json.Unmarshal(httpResponseBytes, rpcResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling: %w", err)
	}
	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("request failed, code: %d, message: %s, data: %s", rpcResponse.Error.Code, rpcResponse.Error.Message, rpcResponse.Error.Data)
	}
	if len(rpcResponse.Result) == 0 {
		return nil, fmt.Errorf("result is empty")
	}
	if err = tmjson.Unmarshal(rpcResponse.Result, result); err != nil {
		return nil, fmt.Errorf("error unmarshalling result: %w", err)
	}
	return result, nil
}

func (c *JsonRpcClient) mapToRequest(method string, params map[string]interface{}) ([]byte, error) {
	var paramsMap = make(map[string]interface{})
	paramsMap["jsonrpc"] = "2.0"
	paramsMap["id"] = 0
	paramsMap["method"] = method
	paramsMap["params"] = params
	return json.Marshal(paramsMap)
}
