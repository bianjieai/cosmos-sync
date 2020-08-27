package service

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgRefundServiceDeposit struct {
		ServiceName string `bson:"service_name" yaml:"service_name"`
		Provider    string `bson:"provider" yaml:"provider"`
		Owner       string `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgRefundServiceDeposit) GetType() string {
	return MsgTypeRefundServiceDeposit
}

func (m *DocMsgRefundServiceDeposit) BuildMsg(v interface{}) {
	msg := v.(MsgRefundServiceDeposit)

	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider.String()
	m.Owner = msg.Owner.String()
}

func (m *DocMsgRefundServiceDeposit) HandleTxMsg(msg MsgRefundServiceDeposit) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Owner, m.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
