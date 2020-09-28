package record

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/cosmos/cosmos-sdk/types"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgRecordCreate).Type():
		docMsg := DocMsgRecordCreate{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
