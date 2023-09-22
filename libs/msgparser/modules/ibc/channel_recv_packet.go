package ibc

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
)

type DocMsgRecvPacket struct {
	Packet      Packet   `bson:"packet"`
	Proof       string   `bson:"proof"`
	ProofHeight int64    `bson:"proof_height"`
	ProofPath   []string `bson:"proof_path"`
	ProofData   string   `bson:"proof_data"`
	ClientID    string   `bson:"client_id"`
	Module      string   `bson:"module"`
	Signer      string   `bson:"signer"`
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	// actual opaque bytes transferred directly to the application module
	Data PacketData `bson:"data"`
	// extended data
	Extra string `bson:"extra"`
}

func (m *DocMsgRecvPacket) GetType() string {
	return MsgTypeRecvPacket
}

func (m *DocMsgRecvPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvPacket)

	m.Proof = string(msg.Proof)
	m.ProofHeight = int64(msg.ProofHeight)
	m.ProofPath = msg.ProofPath
	m.ProofData = string(msg.ProofData)
	m.ClientID = msg.ClientID
	m.Module = msg.Module
	m.Signer = msg.Signer

	m.Packet.Data = DecodeToIBCRecord(msg.Packet)
	m.Packet.Extra = string(msg.Packet.Extra)
}

func (m *DocMsgRecvPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgRecvPacket)
	packetData := DecodeToIBCRecord(msg.Packet)
	addrs = append(addrs, msg.Signer, packetData.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return modules.CreateMsgDocInfo(v, handler)
}
