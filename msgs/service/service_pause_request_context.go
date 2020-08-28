package service

import (
	"encoding/hex"
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type (
	DocMsgPauseRequestContext struct {
		RequestContextID string `bson:"request_context_id" yaml:"request_context_id"`
		Consumer         string `bson:"consumer" yaml:"consumer"`
	}
)

func (m *DocMsgPauseRequestContext) GetType() string {
	return MsgTypePauseRequestContext
}

func (m *DocMsgPauseRequestContext) BuildMsg(v interface{}) {
	msg := v.(MsgPauseRequestContext)

	m.RequestContextID = strings.ToUpper(hex.EncodeToString(msg.RequestContextID))
	m.Consumer = msg.Consumer.String()
}

func (m *DocMsgPauseRequestContext) HandleTxMsg(msg MsgPauseRequestContext) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Consumer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
