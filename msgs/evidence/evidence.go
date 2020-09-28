package evidence
//
//import (
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	. "github.com/bianjieai/irita-sync/msgs"
//)
//
//func HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
//	ok := true
//	switch msgData := msg.(type) {
//	case MsgSubmitEvidence:
//		docMsg := DocMsgSubmitEvidence{}
//		return docMsg.HandleTxMsg(msgData), ok
//	default:
//		ok = false
//	}
//	return MsgDocInfo{}, ok
//}
