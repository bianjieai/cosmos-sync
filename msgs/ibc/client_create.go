package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/bianjieai/irita-sync/models"
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

	m.ClientID = msg.ClientID
	m.Signer = msg.Signer.String()
	m.ClientState = models.Any{TypeUrl: msg.ClientState.GetTypeUrl(), Value: string(msg.ClientState.GetValue())}
	m.ConsensusState = models.Any{TypeUrl: msg.ConsensusState.GetTypeUrl(), Value: string(msg.ConsensusState.GetValue())}

}

func (m *DocMsgCreateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgCreateClient
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
