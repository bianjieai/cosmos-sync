package client

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/ibc"
)

type MsgClient struct {
	Ibc ibc.Client
}

func NewMsgClient() MsgClient {
	return MsgClient{
		Ibc: ibc.NewClient(),
	}
}
