package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgConnectionOpenTry struct {
	ClientId             string       `bson:"client_id"`
	PreviousConnectionId string       `bson:"previous_connection_id"`
	ClientState          string       `bson:"client_state"`
	Counterparty         Counterparty `bson:"counterparty"`
	DelayPeriod          uint64       `bson:"delay_period"`
	CounterpartyVersions []Version    `bson:"counterparty_versions"`
	ProofHeight          Height       `bson:"proof_height"`
	ProofInit            string       `bson:"proof_init"`
	ProofClient          string       `bson:"proof_client"`
	ProofConsensus       string       `bson:"proof_consensus"`
	ConsensusHeight      Height       `bson:"consensus_height"`
	Signer               string       `bson:"signer"`
}

// Counterparty defines the counterparty chain associated with a connection end.
type Counterparty struct {
	ClientId     string       `bson:"client_id" json:"client_id"`
	ConnectionId string       `bson:"connection_id" json:"connection_id"`
	Prefix       MerklePrefix `bson:"prefix" json:"prefix"`
}
type MerklePrefix struct {
	KeyPrefix string `bson:"key_prefix" json:"key_prefix"`
}

type Version struct {
	Identifier string   `bson:"identifier" json:"identifier"`
	Features   []string `bson:"features" json:"features"`
}

func (m *DocMsgConnectionOpenTry) GetType() string {
	return MsgTypeConnectionOpenTry
}

func (m *DocMsgConnectionOpenTry) BuildMsg(v interface{}) {
	msg := v.(*MsgConnectionOpenTry)
	m.Signer = msg.Signer
	m.ClientId = msg.ClientId
	m.PreviousConnectionId = msg.PreviousConnectionId
	m.DelayPeriod = msg.DelayPeriod
	m.ClientState = ConvertAny(msg.ClientState)
	m.ProofInit = utils.MarshalJsonIgnoreErr(msg.ProofInit)
	m.ProofClient = utils.MarshalJsonIgnoreErr(msg.ProofClient)
	m.ProofConsensus = utils.MarshalJsonIgnoreErr(msg.ProofConsensus)
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.ConsensusHeight = loadHeight(msg.ConsensusHeight)
	for _, val := range msg.CounterpartyVersions {
		if val != nil {
			m.CounterpartyVersions = append(m.CounterpartyVersions, Version{
				Identifier: val.Identifier,
				Features:   val.Features,
			})
		}
	}
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(msg.Counterparty), &m.Counterparty)

}

func (m *DocMsgConnectionOpenTry) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgConnectionOpenTry
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
