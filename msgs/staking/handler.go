package staking
//
//import (
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	. "github.com/bianjieai/irita-sync/msgs"
//)
//
//func HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
//	ok := true
//	switch msgData := msg.(type) {
//	case MsgBeginRedelegate:
//		docMsg := DocTxMsgBeginRedelegate{}
//		return docMsg.HandleTxMsg(msgData), ok
//	case MsgUnjail:
//		docMsg := DocTxMsgUnjail{}
//		return docMsg.HandleTxMsg(msgData), ok
//	case MsgStakeBeginUnbonding:
//		docMsg := DocTxMsgBeginUnbonding{}
//		return docMsg.HandleTxMsg(msgData), ok
//	case MsgStakeDelegate:
//		docMsg := DocTxMsgDelegate{}
//		return docMsg.HandleTxMsg(msgData), ok
//	case MsgStakeEdit:
//		docMsg := DocMsgEditValidator{}
//		return docMsg.HandleTxMsg(msgData), ok
//	case MsgStakeCreate:
//		docMsg := DocTxMsgCreateValidator{}
//		return docMsg.HandleTxMsg(msgData), ok
//	default:
//		ok = false
//	}
//	return MsgDocInfo{}, ok
//}
