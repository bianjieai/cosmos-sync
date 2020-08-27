package msg
//import (
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	. "github.com/bianjieai/irita-sync/msgs"
//)
//
//func HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
//	ok := true
//	switch msg.(type) {
//	case MsgAddLiquidity:
//		docMsg := DocTxMsgAddLiquidity{}
//		return docMsg.HandleTxMsg(msg), ok
//	case MsgRemoveLiquidity:
//		docMsg := DocTxMsgRemoveLiquidity{}
//		return docMsg.HandleTxMsg(msg), ok
//	case MsgSwapOrder:
//		docMsg := DocTxMsgSwapOrder{}
//		return docMsg.HandleTxMsg(msg), ok
//	default:
//		ok = false
//	}
//	return MsgDocInfo{}, ok
//}