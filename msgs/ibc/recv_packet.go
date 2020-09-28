package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/msgs/ibc/record"
	"github.com/bianjieai/irita-sync/libs/cdc"
	"gitlab.bianjie.ai/cschain/cschain/modules/ibc/types"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgRecvPacket struct {
	Packet      Packet   `bson:"packet"`
	Proof       string   `bson:"proof"`
	ProofHeight uint64   `bson:"proof_height"`
	ProofPath   []string `bson:"proof_path"`
	ProofData   string   `bson:"proof_data"`
	ClientID    string   `bson:"client_id"`
	Module      string   `bson:"module"`
	Signer      string   `bson:"signer"`
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	// actual opaque bytes transferred directly to the application module
	Data record.IBCRecord `bson:"data"`
	// extended data
	Extra string `bson:"extra"`
}

func (m *DocMsgRecvPacket) GetType() string {
	return MsgTypeRecvPacket
}

func (m *DocMsgRecvPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvPacket)

	m.Proof = string(msg.Proof)
	m.ProofHeight = msg.ProofHeight
	m.ProofPath = msg.ProofPath
	m.ProofData = string(msg.ProofData)
	m.ClientID = msg.ClientID
	m.Module = msg.Module
	m.Signer = msg.Signer.String()

	m.Packet.Data = DecodeToIBCRecord(msg.Packet)
	m.Packet.Extra = string(msg.Packet.Extra)
}
func DecodeToIBCRecord(packet types.Packet) (ibcRecord record.IBCRecord) {
	cdc.GetAmino().UnmarshalJSON([]byte(packet.Data), &ibcRecord)
	return
}

func (m *DocMsgRecvPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRecvPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
