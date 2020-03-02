package msgs

import (
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/types"
)

type (
	DocMsgNFTTransfer struct {
		Sender    string `bson:"sender"`
		Recipient string `bson:"recipient"`
		TokenURI  string `bson:"token_uri"`
		Denom     string `bson:"denom"`
		ID        string `bson:"id"`
	}
)

func (m *DocMsgNFTTransfer) GetType() string {
	return MsgTypeNFTTransfer
}

func (m *DocMsgNFTTransfer) BuildMsg(v interface{}) {
	msg := v.(types.MsgNFTTransfer)

	m.Sender = msg.Sender.String()
	m.Recipient = msg.Recipient.String()
	m.ID = msg.ID
	m.Denom = msg.Denom
	m.TokenURI = msg.TokenURI
}

func (m *DocMsgNFTTransfer) HandleTxMsg(msg types.MsgNFTTransfer) MsgDocInfo {
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
