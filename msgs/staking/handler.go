package staking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.(type) {
	case MsgBeginRedelegate:
		docMsg := DocTxMsgBeginRedelegate{}
		return docMsg.HandleTxMsg(msg), ok
	case MsgUnjail:
		docMsg := DocTxMsgUnjail{}
		return docMsg.HandleTxMsg(msg), ok
	case MsgStakeBeginUnbonding:
		docMsg := DocTxMsgBeginUnbonding{}
		return docMsg.HandleTxMsg(msg), ok
	case MsgStakeDelegate:
		docMsg := DocTxMsgDelegate{}
		return docMsg.HandleTxMsg(msg), ok
	case MsgStakeEdit:
		docMsg := DocTxMsgStakeEdit{}
		return docMsg.HandleTxMsg(msg), ok
	case MsgStakeCreate:
		docMsg := DocTxMsgCreateValidator{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
