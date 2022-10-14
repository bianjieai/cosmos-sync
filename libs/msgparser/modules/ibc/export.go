package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/ibc/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types"
)

type Client interface {
	HandleTxMsg(v types.Msg) (MsgDocInfo, bool)
}
