package msgs

import (
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/types"
)

type (
	DocMsgNFTBurn struct {
		Sender string `json:"sender"`
		ID     string `json:"id"`
		Denom  string `json:"denom"`
	}
)

func (m *DocMsgNFTBurn) GetType() string {
	return MsgTypeNFTBurn
}

func (m *DocMsgNFTBurn) BuildMsg(v interface{}) {
	msg := v.(types.MsgNFTBurn)

	m.Sender = msg.Sender.String()
	m.ID = msg.ID
	m.Denom = msg.Denom
}

func (m *DocMsgNFTBurn) HandleTxMsg(v interface{}) MsgDocInfo {
	var (
		from, to, signer string
		coins            []models.Coin
		docTxMsg         models.DocTxMsg
		complexMsg       bool
		signers          []string
	)

	msg := v.(types.MsgNFTBurn)
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
