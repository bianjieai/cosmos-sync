package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/ibc/types"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
)

type DocMsgChannelOpenConfirm struct {
	PortId      string `bson:"port_id"`
	ChannelId   string `bson:"channel_id"`
	ProofAck    string `bson:"proof_ack"`
	ProofHeight Height `bson:"proof_height"`
	Signer      string `bson:"signer"`
}

func (m *DocMsgChannelOpenConfirm) GetType() string {
	return MsgTypeChannelOpenConfirm
}

func (m *DocMsgChannelOpenConfirm) BuildMsg(v interface{}) {
	msg := v.(*MsgChannelOpenConfirm)
	m.Signer = msg.Signer
	m.PortId = msg.PortId
	m.ChannelId = msg.ChannelId
	m.ProofAck = utils.MarshalJsonIgnoreErr(msg.ProofAck)
	m.ProofHeight = loadHeight(msg.ProofHeight)
}

func (m *DocMsgChannelOpenConfirm) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgChannelOpenConfirm
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
