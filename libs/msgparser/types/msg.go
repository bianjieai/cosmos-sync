package types

import (
	"github.com/okex/exchain/libs/cosmos-sdk/types"
)

type (
	TxMsg struct {
		Type string `bson:"type"`
		Msg  Msg    `bson:"msg"`
	}

	Msg interface {
		GetType() string
		BuildMsg(msg interface{})
	}

	SdkMsg types.Msg
)
