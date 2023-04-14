package client

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/ibc"
	_ "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/uptick_module"
)

type MsgClient struct {
	Ibc ibc.Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Ibc: ibc.NewClient(),
	}
}
