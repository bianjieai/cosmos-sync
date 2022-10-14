package auth

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type Client interface {
	HandleTxMsg(v types.Msg) (MsgDocInfo, bool)
}
