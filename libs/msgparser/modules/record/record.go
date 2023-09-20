package record

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type RecordClient struct {
}

func NewClient() RecordClient {
	return RecordClient{}
}

func (record RecordClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgRecordCreate:
		docMsg := DocMsgRecordCreate{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
