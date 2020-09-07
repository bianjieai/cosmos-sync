package nft

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type DocMsgIssueDenom struct {
	ID     string `bson:"id"`
	Name   string `bson:"name"`
	Sender string `bson:"sender"`
	Schema string `bson:"schema"`
}

func (m *DocMsgIssueDenom) GetType() string {
	return MsgTypeIssueDenom
}

func (m *DocMsgIssueDenom) BuildMsg(v interface{}) {
	msg := v.(*MsgIssueDenom)

	m.Sender = msg.Sender.String()
	m.Schema = msg.Schema
	m.ID = strings.ToLower(msg.ID)
	m.Name = msg.Name
}

func (m *DocMsgIssueDenom) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgIssueDenom
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
