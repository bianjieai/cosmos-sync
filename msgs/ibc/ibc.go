package ibc

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/models"
)

func HandleTxMsg(v types.Msg, timestamp int64) (MsgDocInfo, *models.IbcClient, bool) {
	var (
		msgDocInfo MsgDocInfo
		ibcClient  models.IbcClient
	)
	ok := true
	ibcClient.UpdateAt = timestamp
	switch v.Type() {
	case new(MsgRecvPacket).Type():
		docMsg := DocMsgRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgCreateClient).Type():
		docMsg := DocMsgCreateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		ibcClient.ConsensusState = docMsg.ConsensusState
		ibcClient.ClientState = docMsg.ClientState
		ibcClient.ClientId = docMsg.ClientID
		ibcClient.Signer = docMsg.Signer
		ibcClient.CreateAt = timestamp
	case new(MsgUpdateClient).Type():
		docMsg := DocMsgUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		ibcClient.Header = docMsg.Header
		ibcClient.ClientId = docMsg.ClientID
		ibcClient.Signer = docMsg.Signer
	default:
		ok = false
	}
	return msgDocInfo, &ibcClient, ok
}
