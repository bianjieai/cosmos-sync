package iservice

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgServiceBind struct {
		DefName     string        `bson:"def_name" yaml:"def_name"`
		DefChainID  string        `bson:"def_chain_id" yaml:"def_chain_id"`
		BindChainID string        `bson:"bind_chain_id" yaml:"bind_chain_id"`
		Provider    string        `bson:"provider" yaml:"provider"`
		BindingType string        `bson:"binding_type" yaml:"binding_type"`
		Deposit     []models.Coin `bson:"deposit" yaml:"deposit"`
		Prices      []models.Coin `bson:"price" yaml:"price"`
		Level       Level         `bson:"level" yaml:"level"`
	}

	Level struct {
		AvgRspTime int64 `bson:"avg_rsp_time" yaml:"avg_rsp_time"`
		UsableTime int64 `bson:"usable_time" yaml:"usable_time"`
	}
)

func (m *DocMsgServiceBind) GetType() string {
	return MsgTypeServiceBind
}

func (m *DocMsgServiceBind) BuildMsg(v interface{}) {
	msg := v.(MsgServiceBind)

	m.DefName = msg.DefName
	m.DefChainID = msg.DefChainID
	m.BindChainID = msg.BindChainID
	m.Provider = msg.Provider.String()
	m.BindingType = msg.BindingType.String()
	m.Deposit = models.BuildDocCoins(msg.Deposit)
	m.Prices = models.BuildDocCoins(msg.Prices)
	m.Level = Level{
		AvgRspTime: msg.Level.AvgRspTime,
		UsableTime: msg.Level.UsableTime,
	}
}

func (m *DocMsgServiceBind) HandleTxMsg(msg MsgServiceBind) MsgDocInfo {
	var (
		from, to, signer string
		coins            []models.Coin
		docTxMsg         models.DocTxMsg
		complexMsg       bool
		signers          []string
	)

	from = msg.GetSigners()[0].String()
	to = ""
	coins = models.BuildDocCoins(msg.Deposit)

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
