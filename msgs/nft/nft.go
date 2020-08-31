package nft

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
	case MsgNFTMint:
		docMsg := DocMsgNFTMint{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgNFTEdit:
		docMsg := DocMsgNFTEdit{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgNFTTransfer:
		docMsg := DocMsgNFTTransfer{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgNFTBurn:
		docMsg := DocMsgNFTBurn{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	case MsgIssueDenom:
		docMsg := DocMsgIssueDenom{}
		msgDocInfo = docMsg.HandleTxMsg(msgData)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
