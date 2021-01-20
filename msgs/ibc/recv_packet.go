package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgRecvPacket struct {
	Packet          Packet `bson:"packet"`
	ProofCommitment string `bson:"proof_commitment"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

type Height struct {
	RevisionNumber uint64 `bson:"revision_number"`
	RevisionHeight uint64 `bson:"revision_height"`
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	Sequence           uint64 `bson:"sequence"`
	SourcePort         string `bson:"source_port"`
	SourceChannel      string `bson:"source_channel"`
	DestinationPort    string `bson:"destination_port"`
	DestinationChannel string `bson:"destination_channel"`
	Data               string `bson:"data"`
	TimeoutHeight      Height `bson:"timeout_height"`
	TimeoutTimestamp   uint64 `bson:"timeout_timestamp"`
}

func (m *DocMsgRecvPacket) GetType() string {
	return MsgTypeRecvPacket
}

func (m *DocMsgRecvPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvPacket)
	m.Signer = msg.Signer
	m.ProofHeight = Height{
		RevisionNumber: msg.ProofHeight.RevisionNumber,
		RevisionHeight: msg.ProofHeight.RevisionHeight}
	m.ProofCommitment = utils.MarshalJsonIgnoreErr(m.ProofCommitment)
	m.Packet = Packet{
		Sequence:           msg.Packet.Sequence,
		SourcePort:         msg.Packet.SourcePort,
		SourceChannel:      msg.Packet.SourceChannel,
		DestinationPort:    msg.Packet.DestinationPort,
		DestinationChannel: msg.Packet.DestinationChannel,
		Data:               utils.MarshalJsonIgnoreErr(msg.Packet.Data),
		TimeoutTimestamp:   msg.Packet.TimeoutTimestamp,
		TimeoutHeight: Height{
			RevisionNumber: msg.Packet.TimeoutHeight.RevisionNumber,
			RevisionHeight: msg.Packet.TimeoutHeight.RevisionHeight},
	}

}

func (m *DocMsgRecvPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRecvPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
