package msgs

import (
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/types"
)

type (
	DocMsgNFTMint struct {
		Sender    string `bson:"sender"`
		Recipient string `bson:"recipient"`
		ID        string `bson:"id"`
		Denom     string `bson:"denom"`
		TokenURI  string `bson:"token_uri"`
	}
)

func (m *DocMsgNFTMint) GetType() string {
	return MsgTypeNFTMint
}

func (m *DocMsgNFTMint) BuildMsg(v interface{}) {
	msg := v.(types.MsgNFTMint)

	m.Sender = msg.Sender.String()
	m.Recipient = msg.Recipient.String()
	m.ID = msg.ID
	m.Denom = msg.Denom
	m.TokenURI = msg.TokenURI
}

func (m *DocMsgNFTMint) HandleTxMsg(msg types.MsgNFTMint) MsgDocInfo {
	var (
		from, to, signer string
		coins            []models.Coin
		docTxMsg         models.DocTxMsg
		complexMsg       bool
		signers          []string
	)

	from = msg.Sender.String()
	to = msg.Recipient.String()
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
