package iservice

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msgData := v.(type) {
	case MsgServiceDef:
		docMsg := DocMsgServiceDef{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgServiceBind:
		docMsg := DocMsgServiceBind{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgServiceRequest:
		docMsg := DocMsgServiceRequest{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgServiceResponse:
		docMsg := DocMsgServiceResponse{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
