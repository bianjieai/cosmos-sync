package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type ibcClient struct {
}

func NewClient() Client {
	return ibcClient{}
}

func (ibc ibcClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgRecvPacket:
		docMsg := DocMsgRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgTransfer:
		docMsg := DocMsgTransfer{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgCreateClient:
		docMsg := DocMsgCreateClient{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpdateClient:
		docMsg := DocMsgUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpgradeClient:
		docMsg := DocMsgUpgradeClient{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgSubmitMisbehaviour:
		docMsg := DocMsgSubmitMisbehaviour{}
		msgDocInfo = docMsg.HandleTxMsg(msg)

	case *MsgConnectionOpenInit:
		docMsg := DocMsgConnectionOpenInit{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgConnectionOpenTry:
		docMsg := DocMsgConnectionOpenTry{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgConnectionOpenAck:
		docMsg := DocMsgConnectionOpenAck{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgConnectionOpenConfirm:
		docMsg := DocMsgConnectionOpenConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(msg)

	case *MsgChannelOpenInit:
		docMsg := DocMsgChannelOpenInit{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgChannelOpenTry:
		docMsg := DocMsgChannelOpenTry{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgChannelOpenAck:
		docMsg := DocMsgChannelOpenAck{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgChannelOpenConfirm:
		docMsg := DocMsgChannelOpenConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgChannelCloseInit:
		docMsg := DocMsgChannelCloseInit{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgChannelCloseConfirm:
		docMsg := DocMsgChannelCloseConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgTimeout:
		docMsg := DocMsgTimeout{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgTimeoutOnClose:
		docMsg := DocMsgTimeoutOnClose{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgAcknowledgement:
		docMsg := DocMsgAcknowledgement{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	default:
		ok = false
	}
	return msgDocInfo, ok
}
