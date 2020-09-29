package random

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(msgData sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msgData.Type() {
	case new(MsgRequestRandom).Type():

		txMsg := DocTxMsgRequestRand{}
		return txMsg.HandleTxMsg(msgData), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
