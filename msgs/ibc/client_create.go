package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"gitlab.bianjie.ai/cschain/cschain/modules/ibc/core/types"
	"encoding/json"
)

// MsgCreateClient defines a message to create an IBC client
type DocMsgCreateClient struct {
	ClientID       string      `bson:"client_id" yaml:"client_id"`
	ClientState    interface{} `bson:"client_state"`
	ConsensusState interface{} `bson:"consensus_state"`
	Signer         string      `bson:"signer" yaml:"signer"`
}

func (m *DocMsgCreateClient) GetType() string {
	return MsgTypeCreateClient
}

func (m *DocMsgCreateClient) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateClient)

	if clientState, err := types.UnpackClientState(msg.ClientState); err == nil {
		data, _ := json.Marshal(clientState)
		m.ClientState = string(data)
	}

	if consensusState, err := types.UnpackConsensusState(msg.ConsensusState); err == nil {
		data, _ := json.Marshal(consensusState)
		m.ConsensusState = string(data)
	}

	m.ClientID = msg.ClientID
	m.Signer = msg.Signer

}

func (m *DocMsgCreateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
	)

	return CreateMsgDocInfo(v, func() (Msg, []string) {
		return m, addrs
	})
}
