package token

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
	case MsgMintToken:
		docMsg := DocMsgMintToken{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgEditToken:
		docMsg := DocMsgEditToken{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgIssueToken:
		docMsg := DocMsgIssueToken{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgTransferTokenOwner:
		docMsg := DocMsgTransferTokenOwner{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
