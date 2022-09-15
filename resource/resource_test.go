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
	data, err := HttpGet("https://raw.githubusercontent.com/cosmos/chain-registry/master/sentinel/chain.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	nodeurl, err := LoadRpcResource(string(data))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(nodeurl)
}
