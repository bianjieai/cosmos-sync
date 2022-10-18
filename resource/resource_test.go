package resource

import "testing"

func TestHttpGet(t *testing.T) {
	data, err := HttpGet("https://raw.githubusercontent.com/cosmos/chain-registry/master/sentinel/chain.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(data))
}

func TestLoadRpcResource(t *testing.T) {
	data, err := HttpGet("https://raw.githubusercontent.com/cosmos/chain-registry/master/likecoin/chain.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	nodeurl, err := LoadRpcResource(string(data), "likecoin_mainnet_2")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(nodeurl)
}

func TestHandleUri(t *testing.T) {
	t.Log(HandleUri("https://mainnet-full-node-02.hosts.cudos.org:36657"))
	t.Log(HandleUri("https://mainnet-node.like.co/rpc/"))
}
