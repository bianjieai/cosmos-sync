package service

import (
	"encoding/hex"
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type (
	DocMsgStartRequestContext struct {
		RequestContextID string `bson:"request_context_id" yaml:"request_context_id"`
		Consumer         string `bson:"consumer" yaml:"consumer"`
	}
)

func (m *DocMsgStartRequestContext) GetType() string {
	return MsgTypeStartRequestContext
}

func (m *DocMsgStartRequestContext) BuildMsg(v interface{}) {
	msg := v.(MsgStartRequestContext)

	m.RequestContextID = strings.ToUpper(hex.EncodeToString(msg.RequestContextID))
	m.Consumer = msg.Consumer.String()
}

func (m *DocMsgStartRequestContext) HandleTxMsg(msg MsgStartRequestContext) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Consumer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
