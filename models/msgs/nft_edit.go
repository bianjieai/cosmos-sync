package msgs

import (
	"gitlab.bianjie.ai/irita/ex-sync/models"
	"gitlab.bianjie.ai/irita/ex-sync/types"
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

func (m *DocMsgNFTEdit) HandleTxMsg(v interface{}) MsgDocInfo {
	var (
		from, to, signer string
		coins            []models.Coin
		docTxMsg         models.DocTxMsg
		complexMsg       bool
		signers          []string
	)

	msg := v.(types.MsgNFTEdit)
	from = msg.Sender.String()
	to = ""
	coins = models.BuildDocCoins(nil)

	m.BuildMsg(v)
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
