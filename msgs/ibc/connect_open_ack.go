package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgConnectionOpenAck struct {
	ConnectionId             string  `bson:"connection_id"`
	CounterpartyConnectionId string  `bson:"counterparty_connection_id"`
	Version                  Version `bson:"version"`
	ClientState              string  `bson:"client_state"`
	ProofHeight              Height  `bson:"proof_height"`
	ProofTry                 string  `bson:"proof_try"`
	ProofClient              string  `bson:"proof_client"`
	ProofConsensus           string  `bson:"proof_consensus"`
	ConsensusHeight          Height  `bson:"consensus_height"`
	Signer                   string  `bson:"signer"`
}

func (m *DocMsgConnectionOpenAck) GetType() string {
	return MsgTypeConnectionOpenAck
}

func (m *DocMsgConnectionOpenAck) BuildMsg(v interface{}) {
	msg := v.(*MsgConnectionOpenAck)
	m.Signer = msg.Signer
	m.ConnectionId = msg.ConnectionId
	m.CounterpartyConnectionId = msg.CounterpartyConnectionId
	m.ClientState = ConvertAny(msg.ClientState)
	if msg.Version != nil {
		m.Version = Version{
			Identifier: msg.Version.Identifier,
			Features:   msg.Version.Features,
		}
	}
	m.ProofTry = utils.MarshalJsonIgnoreErr(msg.ProofTry)
	m.ProofClient = utils.MarshalJsonIgnoreErr(msg.ProofClient)
	m.ProofConsensus = utils.MarshalJsonIgnoreErr(msg.ProofConsensus)
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.ConsensusHeight = loadHeight(msg.ConsensusHeight)

}

func (m *DocMsgConnectionOpenAck) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgConnectionOpenAck
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
