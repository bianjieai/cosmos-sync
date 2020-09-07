package token

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgIssueToken struct {
	Symbol        string `bson:"symbol"`
	Name          string `bson:"name"`
	Scale         uint32 `bson:"scale"`
	MinUnit       string `bson:"min_unit"`
	InitialSupply uint64 `bson:"initial_supply"`
	MaxSupply     uint64 `bson:"max_supply"`
	Mintable      bool   `bson:"mintable"`
	Owner         string `bson:"owner"`
}

func (m *DocMsgIssueToken) GetType() string {
	return MsgTypeIssueToken
}

func (m *DocMsgIssueToken) BuildMsg(v interface{}) {
	msg := v.(*MsgIssueToken)

	m.Symbol = msg.Symbol
	m.Name = msg.Name
	m.Scale = msg.Scale
	m.MinUnit = msg.MinUnit
	m.InitialSupply = msg.InitialSupply
	m.Owner = msg.Owner.String()
	m.MaxSupply = msg.MaxSupply
	m.Mintable = msg.Mintable
}

func (m *DocMsgIssueToken) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgIssueToken
	)

	ConvertMsg(v,&msg)
	addrs = append(addrs, msg.Owner.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
