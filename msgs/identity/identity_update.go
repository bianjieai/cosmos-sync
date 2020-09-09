package identity

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/models"
)

// MsgUpdateIdentity defines a message to update an identity
type DocMsgUpdateIdentity struct {
	ID          string      `bson:"id"`
	PubKey      *PubKeyInfo `bson:"pubkey"`
	Certificate string      `bson:"certificate"`
	Credentials string      `bson:"credentials"`
	Owner       string      `bson:"owner"`
}

func (m *DocMsgUpdateIdentity) GetType() string {
	return MsgTypeUpdateIdentity
}

func (m *DocMsgUpdateIdentity) BuildMsg(v interface{}) {
	msg := v.(MsgUpdateIdentity)
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

func (m *DocMsgUpdateIdentity) HandleTxMsg(msg MsgUpdateIdentity) MsgDocInfo {
	var (
		docTxMsg models.DocTxMsg
		signers  []string
		addrs    []string
	)

	_, signers = models.BuildDocSigners(msg.GetSigners())
	addrs = append(addrs, signers...)

	m.BuildMsg(msg)
	docTxMsg = models.DocTxMsg{
		Type: m.GetType(),
		Msg:  m,
	}
	addrs = append(addrs, m.Owner)

	return MsgDocInfo{
		DocTxMsg: docTxMsg,
		Signers:  signers,
		Addrs:    addrs,
	}
}
