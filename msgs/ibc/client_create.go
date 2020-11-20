package ibc

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

// MsgCreateClient defines a message to create an IBC client
type DocMsgCreateClient struct {
	ClientID       string     `bson:"client_id" yaml:"client_id"`
	ClientState    models.Any `bson:"client_state"`
	ConsensusState models.Any `bson:"consensus_state"`
	Signer         string     `bson:"signer" yaml:"signer"`
}

func (m *DocMsgCreateClient) GetType() string {
	return MsgTypeCreateClient
}

func (m *DocMsgCreateClient) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateClient)

	m.ClientID = msg.ClientId
	m.Signer = msg.Signer
	m.ClientState = models.Any{TypeUrl: msg.ClientState.GetTypeUrl(), Value: string(msg.ClientState.GetValue())}
	m.ConsensusState = models.Any{TypeUrl: msg.ConsensusState.GetTypeUrl(), Value: string(msg.ConsensusState.GetValue())}

}

func (m *DocMsgCreateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
	)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}