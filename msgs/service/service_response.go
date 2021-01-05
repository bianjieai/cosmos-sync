package service

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgServiceResponse struct {
		RequestID string `bson:"request_id" yaml:"request_id"`
		Provider  string `bson:"provider" yaml:"provider"`
		Output    string `bson:"output" yaml:"output"`
		Result    string `bson:"result"`
	}
)

func (m *DocMsgServiceResponse) GetType() string {
	return MsgTypeRespondService
}

func (m *DocMsgServiceResponse) BuildMsg(msg interface{}) {
	v := msg.(*MsgRespondService)

	m.RequestID = v.RequestId
	m.Provider = v.Provider
	//m.Output = hex.EncodeToString(v.Output)
	m.Output = v.Output
	m.Result = v.Result
}

func (m *DocMsgServiceResponse) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRespondService
	)
	ConvertMsg(v, &msg)

	addrs = append(addrs, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
