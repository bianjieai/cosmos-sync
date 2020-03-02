package msgs

import (
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/types"
)

type (
	DocMsgNFTEdit struct {
		Sender   string `bson:"sender"`
		ID       string `bson:"id"`
		Denom    string `bson:"denom"`
		TokenURI string `bson:"token_uri"`
	}
)

func (m *DocMsgNFTEdit) GetType() string {
	return MsgTypeNFTEdit
}

func (m *DocMsgNFTEdit) BuildMsg(v interface{}) {
	msg := v.(types.MsgNFTEdit)

	m.Sender = msg.Sender.String()
	m.ID = msg.ID
	m.Denom = msg.Denom
	m.TokenURI = msg.TokenURI
}

func (m *DocMsgNFTEdit) HandleTxMsg(msg types.MsgNFTEdit) MsgDocInfo {
	var (
		from, to, signer string
		coins            []models.Coin
		docTxMsg         models.DocTxMsg
		complexMsg       bool
		signers          []string
	)

	from = msg.Sender.String()
	to = ""
	coins = models.BuildDocCoins(nil)

	m.BuildMsg(msg)
	docTxMsg = models.DocTxMsg{
		Type: m.GetType(),
		Msg:  m,
	}
	complexMsg = false

	signer, signers = models.BuildDocSigners(msg.GetSigners())

	return MsgDocInfo{
		From:       from,
		To:         to,
		Coins:      coins,
		Signer:     signer,
		DocTxMsg:   docTxMsg,
		ComplexMsg: complexMsg,
		Signers:    signers,
	}
}
