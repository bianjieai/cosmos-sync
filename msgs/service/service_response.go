package service

import (
	"encoding/hex"
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
	v := msg.(MsgRespondService)

	m.RequestID = hex.EncodeToString(v.RequestID)
	m.Provider = v.Provider.String()
	//m.Output = hex.EncodeToString(v.Output)
	m.Output = v.Output
	m.Result = v.Result
}

func (m *DocMsgServiceResponse) HandleTxMsg(msg MsgRespondService) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, msg.Provider.String(), msg.Provider.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
