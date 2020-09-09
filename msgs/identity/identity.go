package identity

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)
func HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo  MsgDocInfo
	)
	ok := true
	switch msgData := v.(type) {
	case MsgCreateIdentity:
		docMsg := DocMsgCreateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgUpdateIdentity:
		docMsg := DocMsgUpdateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
