package identity

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
)

// DocMsgUpdateIdentity MsgUpdateIdentity defines a message to update an identity
type DocMsgUpdateIdentity struct {
	Id          string      `bson:"id"`
	PubKey      *PubKeyInfo `bson:"pubkey"`
	Certificate string      `bson:"certificate"`
	Credentials string      `bson:"credentials"`
	Owner       string      `bson:"owner"`
	ExTemporary ExTemporary `bson:"ex"`
}

func (m *DocMsgUpdateIdentity) GetType() string {
	return MsgTypeUpdateIdentity
}

func (m *DocMsgUpdateIdentity) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateIdentity)
	m.Id = msg.Id
	m.Owner = msg.Owner
	if msg.PubKey != nil {
		m.PubKey = &PubKeyInfo{
			PubKey:    msg.PubKey.PubKey,
			Algorithm: int32(msg.PubKey.Algorithm),
		}
	}
	m.Certificate = msg.Certificate
	m.Credentials = msg.Credentials

	if m.Certificate != "" {
		m.ExTemporary = ExTemporary{
			CertPubKey: getPubKeyFromCertificate(m.Certificate),
		}
	}
}

func (m *DocMsgUpdateIdentity) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUpdateIdentity)
	addrs = append(addrs, msg.Owner)
	addrs = append(addrs, msg.Owner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return modules.CreateMsgDocInfo(v, handler)
}
