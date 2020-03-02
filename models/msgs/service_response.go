package msgs

import (
	"encoding/hex"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/types"
)

type (
	DocMsgServiceResponse struct {
		ReqChainID string `bson:"req_chain_id" yaml:"req_chain_id"`
		RequestID  string `bson:"request_id" yaml:"request_id"`
		Provider   string `bson:"provider" yaml:"provider"`
		Output     string `bson:"output" yaml:"output"`
		ErrorMsg   string `bson:"error_msg" yaml:"error_msg"`
	}
)

func (m *DocMsgServiceResponse) GetType() string {
	return MsgTypeServiceResponse
}

func (m *DocMsgServiceResponse) BuildMsg(msg interface{}) {
	v := msg.(types.MsgServiceResponse)

	m.ReqChainID = v.ReqChainID
	m.RequestID = v.RequestID
	m.Provider = v.Provider.String()
	m.Output = hex.EncodeToString(v.Output)
	m.ErrorMsg = hex.EncodeToString(v.ErrorMsg)
}

func (m *DocMsgServiceResponse) HandleTxMsg(msg types.MsgServiceResponse) MsgDocInfo {
	var (
		from, to, signer string
		coins            []models.Coin
		docTxMsg         models.DocTxMsg
		complexMsg       bool
		signers          []string
	)

	from = msg.Provider.String()
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
