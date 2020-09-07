package service

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(msg types.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.Type() {
	case new(MsgDefineService).Type():
		docMsg := DocMsgDefineService{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgBindService).Type():
		docMsg := DocMsgBindService{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgCallService).Type():
		docMsg := DocMsgCallService{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgRespondService).Type():
		docMsg := DocMsgServiceResponse{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgUpdateServiceBinding).Type():
		docMsg := DocMsgUpdateServiceBinding{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgSetWithdrawAddress).Type():
		docMsg := DocMsgSetWithdrawAddress{}
		msgData := MsgSetWithdrawAddress{}
		ConvertMsg(msg, &msgData)
		if msgData.Owner.String() == "" {
			return MsgDocInfo{}, false
		}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgDisableServiceBinding).Type():
		docMsg := DocMsgDisableServiceBinding{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgEnableServiceBinding).Type():
		docMsg := DocMsgEnableServiceBinding{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgRefundServiceDeposit).Type():
		docMsg := DocMsgRefundServiceDeposit{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgUpdateRequestContext).Type():
		docMsg := DocMsgUpdateRequestContext{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgPauseRequestContext).Type():
		docMsg := DocMsgPauseRequestContext{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgStartRequestContext).Type():
		docMsg := DocMsgStartRequestContext{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgKillRequestContext).Type():
		docMsg := DocMsgKillRequestContext{}
		return docMsg.HandleTxMsg(msg), ok

	case new(MsgWithdrawEarnedFees).Type():
		docMsg := DocMsgWithdrawEarnedFees{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
