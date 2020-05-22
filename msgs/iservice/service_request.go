package iservice

import (
	"encoding/hex"
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgServiceRequest struct {
		DefChainID  string        `bson:"def_chain_id" yaml:"def_chain_id"`
		DefName     string        `bson:"def_name" yaml:"def_name"`
		BindChainID string        `bson:"bind_chain_id" yaml:"bind_chain_id"`
		ReqChainID  string        `bson:"req_chain_id" yaml:"req_chain_id"`
		MethodID    int16         `bson:"method_id" yaml:"method_id"`
		Provider    string        `bson:"provider" yaml:"provider"`
		Consumer    string        `bson:"consumer" yaml:"consumer"`
		Input       string        `bson:"input" yaml:"input"`
		ServiceFee  []models.Coin `bson:"service_fee" yaml:"service_fee"`
		Profiling   bool          `bson:"profiling" yaml:"profiling"`
	}
)

func (m *DocMsgServiceRequest) GetType() string {
	return MsgTypeServiceRequest
}

func (m *DocMsgServiceRequest) BuildMsg(msg interface{}) {
	v := msg.(MsgServiceRequest)

	m.DefChainID = v.DefChainID
	m.DefName = v.DefName
	m.BindChainID = v.BindChainID
	m.ReqChainID = v.ReqChainID
	m.MethodID = v.MethodID
	m.Provider = v.Provider.String()
	m.Consumer = v.Consumer.String()
	m.Input = hex.EncodeToString(v.Input)
	m.ServiceFee = models.BuildDocCoins(v.ServiceFee)
	m.Profiling = v.Profiling
}

func (m *DocMsgServiceRequest) HandleTxMsg(msg MsgServiceRequest) MsgDocInfo {
	var (
		from, to, signer string
		coins            []models.Coin
		docTxMsg         models.DocTxMsg
		complexMsg       bool
		signers          []string
	)

	from = msg.Consumer.String()
	to = msg.Provider.String()
	coins = models.BuildDocCoins(msg.ServiceFee)

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
