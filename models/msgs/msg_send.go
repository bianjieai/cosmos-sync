package msgs

import (
	"gitlab.bianjie.ai/irita/ex-sync/models"
	"gitlab.bianjie.ai/irita/ex-sync/types"
)

type (
	DocMsgSend struct {
		FromAddress string        `json:"from_address"`
		ToAddress   string        `json:"to_address"`
		Amount      []models.Coin `json:"amount"`
	}
)

func (m *DocMsgSend) Type() string {
	return MsgTypeSend
}

func (m *DocMsgSend) BuildMsg(v interface{}) {
	msg := v.(types.MsgSend)
	m.FromAddress = msg.FromAddress.String()
	m.ToAddress = msg.ToAddress.String()
	m.Amount = models.BuildDocCoins(msg.Amount)
}

func (m *DocMsgSend) HandleTxMsg(v interface{}) MsgDocInfo {
	var (
		complexMsg       bool
		from, to, signer string
		signers          []string
		amount           []models.Coin
		docTxMsg         models.DocTxMsg
	)
	msg := v.(types.MsgSend)

	from = msg.FromAddress.String()
	to = msg.ToAddress.String()
	signer, signers = models.BuildDocSigners(msg.GetSigners())
	amount = models.BuildDocCoins(msg.Amount)

	m.BuildMsg(v)
	docTxMsg = models.DocTxMsg{
		Type: m.Type(),
		Msg:  m,
	}

	return MsgDocInfo{
		ComplexMsg: complexMsg,
		From:       from,
		To:         to,
		Coins:      amount,
		Signer:     signer,
		Signers:    signers,
		DocTxMsg:   docTxMsg,
	}
}
