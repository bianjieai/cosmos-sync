package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgUpgradeClient struct {
	ClientId                   string `bson:"client_id"`
	ClientState                string `bson:"client_state"`
	ConsensusState             string `bson:"consensus_state"`
	ProofUpgradeClient         string `bson:"proof_upgrade_client"`
	ProofUpgradeConsensusState string `bson:"proof_upgrade_consensus_state"`
	Signer                     string `bson:"signer"`
}

func (m *DocMsgUpgradeClient) GetType() string {
	return MsgTypeUpgradeClient
}

func (m *DocMsgUpgradeClient) BuildMsg(v interface{}) {
	msg := v.(*MsgUpgradeClient)
	m.Signer = msg.Signer
	m.ClientId = msg.ClientId
	m.ClientState = ConvertAny(msg.ClientState)
	m.ConsensusState = ConvertAny(msg.ConsensusState)
	m.ProofUpgradeClient = utils.MarshalJsonIgnoreErr(msg.ProofUpgradeClient)
	m.ProofUpgradeConsensusState = utils.MarshalJsonIgnoreErr(msg.ProofUpgradeConsensusState)
}

func (m *DocMsgUpgradeClient) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgUpgradeClient
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
