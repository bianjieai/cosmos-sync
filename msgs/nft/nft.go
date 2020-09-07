package nft

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)


func HandleTxMsg(msg types.Msg) (MsgDocInfo, bool) {

	switch  msg.Type() {
	case new(MsgNFTMint).Type():
		docMsg := DocMsgNFTMint{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgNFTEdit).Type():
		docMsg := DocMsgNFTEdit{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgNFTTransfer).Type():
		docMsg := DocMsgNFTTransfer{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgNFTBurn).Type():
		docMsg := DocMsgNFTBurn{}
		return docMsg.HandleTxMsg(msg), true
	case new(MsgIssueDenom).Type():
		docMsg := DocMsgIssueDenom{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
