package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type IbcClient struct {
}

func NewClient() IbcClient {
	return IbcClient{}
}

func (ibc IbcClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {

	switch msg := v.(type) {
	case *MsgRecvPacket:
		docMsg := DocMsgRecvPacket{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgCreateClient:
		docMsg := DocMsgCreateClient{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUpdateClient:
		docMsg := DocMsgUpdateClient{}
		return docMsg.HandleTxMsg(msg), true
	default:
		return MsgDocInfo{}, false
	}

}
