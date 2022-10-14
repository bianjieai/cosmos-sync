package exchain_ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	exchainTypes "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/exchain-ibc/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types"
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
	case *exchainTypes.MsgRecvPacket:
		docMsg := DocMsgRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *exchainTypes.MsgTransfer:
		docMsg := DocMsgTransfer{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgCreateClient:
	//	docMsg := DocMsgCreateClient{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	case *exchainTypes.MsgUpdateClient:
		docMsg := DocMsgUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgUpgradeClient:
	//	docMsg := DocMsgUpgradeClient{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgSubmitMisbehaviour:
	//	docMsg := DocMsgSubmitMisbehaviour{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)

	//case *exchainTypes.MsgConnectionOpenInit:
	//	docMsg := DocMsgConnectionOpenInit{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgConnectionOpenTry:
	//	docMsg := DocMsgConnectionOpenTry{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgConnectionOpenAck:
	//	docMsg := DocMsgConnectionOpenAck{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgConnectionOpenConfirm:
	//	docMsg := DocMsgConnectionOpenConfirm{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)

	//case *exchainTypes.MsgChannelOpenInit:
	//	docMsg := DocMsgChannelOpenInit{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgChannelOpenTry:
	//	docMsg := DocMsgChannelOpenTry{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgChannelOpenAck:
	//	docMsg := DocMsgChannelOpenAck{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	case *exchainTypes.MsgChannelOpenConfirm:
		docMsg := DocMsgChannelOpenConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgChannelCloseInit:
	//	docMsg := DocMsgChannelCloseInit{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgChannelCloseConfirm:
	//	docMsg := DocMsgChannelCloseConfirm{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	case *exchainTypes.MsgTimeout:
		docMsg := DocMsgTimeout{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	//case *exchainTypes.MsgTimeoutOnClose:
	//	docMsg := DocMsgTimeoutOnClose{}
	//	msgDocInfo = docMsg.HandleTxMsg(msg)
	case *exchainTypes.MsgAcknowledgement:
		docMsg := DocMsgAcknowledgement{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	default:
		ok = false
	}
	return msgDocInfo, ok
}
