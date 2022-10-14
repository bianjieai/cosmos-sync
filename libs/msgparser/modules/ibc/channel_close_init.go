package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
)

type DocMsgChannelCloseInit struct {
	PortId    string `bson:"port_id"`
	ChannelId string `bson:"channel_id"`
	Signer    string `bson:"signer"`
}

func (m *DocMsgChannelCloseInit) GetType() string {
	return MsgTypeChannelCloseInit
}

func (m *DocMsgChannelCloseInit) BuildMsg(v interface{}) {
	msg := v.(*MsgChannelCloseInit)
	m.Signer = msg.Signer
	m.PortId = msg.PortId
	m.ChannelId = msg.ChannelId
}

func (m *DocMsgChannelCloseInit) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgChannelCloseInit
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
