package identity

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

// PubKey represents a public key along with the corresponding algorithm
type PubKeyInfo struct {
	PubKey    string `bson:"pubkey"`
	Algorithm int32  `bson:"algorithm"`
}

// MsgCreateIdentity defines a message to create an identity
type DocMsgCreateIdentity struct {
	ID          string      `bson:"id"`
	PubKey      *PubKeyInfo `bson:"pubkey"`
	Certificate string      `bson:"certificate"`
	Credentials string      `bson:"credentials"`
	Owner       string      `bson:"owner"`
}

func (m *DocMsgCreateIdentity) GetType() string {
	return MsgTypeCreateIdentity
}

func (m *DocMsgCreateIdentity) BuildMsg(v interface{}) {
	msg := v.(MsgCreateIdentity)
	m.ID = msg.ID.String()
	m.Owner = msg.Owner.String()
	if msg.PubKey != nil {
		m.PubKey = &PubKeyInfo{
			PubKey:    msg.PubKey.PubKey.String(),
			Algorithm: int32(msg.PubKey.Algorithm),
		}
	}
	m.Certificate = msg.Certificate
	m.Credentials = msg.Credentials
}

func (m *DocMsgCreateIdentity) HandleTxMsg(msg MsgCreateIdentity) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, msg.Owner.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
