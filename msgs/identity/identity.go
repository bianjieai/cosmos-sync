package identity

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgCreateIdentity).Type():
		docMsg := DocMsgCreateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	case new(MsgUpdateIdentity).Type():
		docMsg := DocMsgUpdateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
