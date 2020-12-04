package identity

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type (
	// PubKey represents a public key along with the corresponding algorithm
	PubKeyInfo struct {
		PubKey    string `bson:"pubkey"`
		Algorithm int32  `bson:"algorithm"`
	}

	ExTemporary struct {
		CertPubKey PubKeyInfo `bson:"cert_pub_key"`
	}
)

// MsgCreateIdentity defines a message to create an identity
type DocMsgCreateIdentity struct {
	Id          string      `bson:"id"`
	PubKey      *PubKeyInfo `bson:"pubkey"`
	Certificate string      `bson:"certificate"`
	Credentials string      `bson:"credentials"`
	Owner       string      `bson:"owner"`
	ExTemporary ExTemporary `bson:"ex"`
}

func (m *DocMsgCreateIdentity) GetType() string {
	return MsgTypeCreateIdentity
}

func (m *DocMsgCreateIdentity) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateIdentity)
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

func (m *DocMsgCreateIdentity) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg MsgCreateIdentity
	)
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Owner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
