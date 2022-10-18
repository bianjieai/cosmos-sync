package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
)

type DocMsgChannelOpenInit struct {
	PortId  string  `bson:"port_id"`
	Channel Channel `bson:"channel"`
	Signer  string  `bson:"signer"`
}

func (m *DocMsgChannelOpenInit) GetType() string {
	return MsgTypeChannelOpenInit
}

func (m *DocMsgChannelOpenInit) BuildMsg(v interface{}) {
	msg := v.(*MsgChannelOpenInit)
	m.Signer = msg.Signer
	m.PortId = msg.PortId
	m.Channel = loadChannel(msg.Channel)
}

func (m *DocMsgChannelOpenInit) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgChannelOpenInit
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
