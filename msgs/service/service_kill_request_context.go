package service

import (
	"encoding/hex"
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type (
	DocMsgKillRequestContext struct {
		RequestContextID string `bson:"request_context_id" yaml:"request_context_id"`
		Consumer         string `bson:"consumer" yaml:"consumer"`
	}
)

func (m *DocMsgKillRequestContext) GetType() string {
	return MsgTypeKillRequestContext
}

func (m *DocMsgKillRequestContext) BuildMsg(v interface{}) {
	msg := v.(MsgKillRequestContext)

	m.RequestContextID = strings.ToUpper(hex.EncodeToString(msg.RequestContextID))
	m.Consumer = msg.Consumer.String()
}

func (m *DocMsgKillRequestContext) HandleTxMsg(msg MsgKillRequestContext) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, msg.Consumer.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
