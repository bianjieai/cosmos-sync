package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.(type) {
	case MsgSubmitProposal:
		docMsg := DocTxMsgSubmitProposal{}
		return docMsg.HandleTxMsg(msg), ok
	case MsgVote:
		docMsg := DocTxMsgVote{}
		return docMsg.HandleTxMsg(msg), ok
	case MsgDeposit:
		docMsg := DocTxMsgDeposit{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
