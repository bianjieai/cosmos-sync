package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

// MsgCreateClient defines a message to create an IBC client
type DocMsgCreateClient struct {
	ClientState    string `bson:"client_state"`
	ConsensusState string `bson:"consensus_state"`
	Signer         string `bson:"signer" yaml:"signer"`
}

func (m *DocMsgCreateClient) GetType() string {
	return MsgTypeCreateClient
}

func (m *DocMsgCreateClient) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateClient)

	m.Signer = msg.Signer
	m.ClientState = ConvertAny(msg.ClientState)
	m.ConsensusState = ConvertAny(msg.ConsensusState)
}

func (m *DocMsgCreateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgCreateClient
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
