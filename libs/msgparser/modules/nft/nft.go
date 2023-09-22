package nft

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type NftClient struct {
}

func NewClient() NftClient {
	return NftClient{}
}

func (nft NftClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {

	switch msg := v.(type) {
	case *MsgNFTMint:
		docMsg := DocMsgNFTMint{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgNFTEdit:
		docMsg := DocMsgNFTEdit{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgNFTTransfer:
		docMsg := DocMsgNFTTransfer{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgNFTBurn:
		docMsg := DocMsgNFTBurn{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgIssueDenom:
		docMsg := DocMsgIssueDenom{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgTransferDenom:
		docMsg := DocMsgTransferDenom{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
