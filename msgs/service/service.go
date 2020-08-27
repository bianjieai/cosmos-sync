package service

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
	case MsgDefineService:
		docMsg := DocMsgDefineService{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgBindService:
		docMsg := DocMsgBindService{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgCallService:
		docMsg := DocMsgCallService{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgRespondService:
		docMsg := DocMsgServiceResponse{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgUpdateServiceBinding:
		docMsg := DocMsgUpdateServiceBinding{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgSetWithdrawAddress:
		docMsg := DocMsgSetWithdrawAddress{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgDisableServiceBinding:
		docMsg := DocMsgDisableServiceBinding{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgEnableServiceBinding:
		docMsg := DocMsgEnableServiceBinding{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgRefundServiceDeposit:
		docMsg := DocMsgRefundServiceDeposit{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgUpdateRequestContext:
		docMsg := DocMsgUpdateRequestContext{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgPauseRequestContext:
		docMsg := DocMsgPauseRequestContext{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgStartRequestContext:
		docMsg := DocMsgStartRequestContext{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgKillRequestContext:
		docMsg := DocMsgKillRequestContext{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgWithdrawEarnedFees:
		docMsg := DocMsgWithdrawEarnedFees{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
