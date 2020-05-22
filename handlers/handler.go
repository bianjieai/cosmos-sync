package handlers

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/msgs/iservice"
	"github.com/bianjieai/irita-sync/msgs/nft"
	"github.com/bianjieai/irita-sync/msgs/base"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo) {
	if iServiceDocinfo, ok := iservice.HandleTxMsg(v); ok {
		return iServiceDocinfo
	}
	if nftDocinfo, ok := nft.HandleTxMsg(v); ok {
		return nftDocinfo
	}
	if baseDocinfo, ok := base.HandleTxMsg(v); ok {
		return baseDocinfo
	}
	return MsgDocInfo{}
}
