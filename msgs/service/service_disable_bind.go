package service

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgDisableServiceBinding struct {
		ServiceName string `bson:"service_name" yaml:"service_name"`
		Provider    string `bson:"provider" yaml:"provider"`
		Owner       string `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgDisableServiceBinding) GetType() string {
	return MsgTypeDisableServiceBinding
}

func (m *DocMsgDisableServiceBinding) BuildMsg(v interface{}) {
	msg := v.(*MsgDisableServiceBinding)

	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider.String()
	m.Owner = msg.Owner.String()
}

func (m *DocMsgDisableServiceBinding) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg MsgDisableServiceBinding
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Owner.String(), msg.Provider.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
