package gov
//
//import (
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	. "github.com/bianjieai/irita-sync/msgs"
//)
//
//func HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
//	ok := true
//	switch msgData := msg.(type) {
//	case MsgSubmitProposal:
//		docMsg := DocTxMsgSubmitProposal{}
//		return docMsg.HandleTxMsg(msgData), ok
//	case MsgVote:
//		docMsg := DocTxMsgVote{}
//		return docMsg.HandleTxMsg(msgData), ok
//	case MsgDeposit:
//		docMsg := DocTxMsgDeposit{}
//		return docMsg.HandleTxMsg(msgData), ok
//	default:
//		ok = false
//	}
//	return MsgDocInfo{}, ok
//}
