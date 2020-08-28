package msg

//import (
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	. "github.com/bianjieai/irita-sync/msgs"
//)
//
//func HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
//	ok := true
//	switch msg.(type) {
//	case MsgClaimHTLC:
//		docMsg := DocTxMsgClaimHTLC{}
//		return docMsg.HandleTxMsg(msg), ok
//	case MsgCreateHTLC:
//		docMsg := DocTxMsgCreateHTLC{}
//		return docMsg.HandleTxMsg(msg), ok
//	case MsgRefundHTLC:
//		docMsg := DocTxMsgRefundHTLC{}
//		return docMsg.HandleTxMsg(msg), ok
//	default:
//		ok = false
//	}
//	return MsgDocInfo{}, ok
//}
