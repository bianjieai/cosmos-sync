package token

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgEditToken struct {
	Symbol    string `bson:"symbol"`
	Name      string `bson:"name"`
	MaxSupply uint64 `bson:"max_supply"`
	Mintable  bool   `bson:"mintable"`
	Owner     string `bson:"owner"`
}

func (m *DocMsgEditToken) GetType() string {
	return MsgTypeEditToken
}

func (m *DocMsgEditToken) BuildMsg(v interface{}) {
	msg := v.(MsgEditToken)

	m.Symbol = msg.Symbol
	m.Owner = msg.Owner.String()
	m.Name = msg.Name
	m.MaxSupply = msg.MaxSupply
	m.Mintable = msg.Mintable.ToBool()
}

func (m *DocMsgEditToken) HandleTxMsg(msg MsgEditToken) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, msg.Owner.String())
	handler := func() (Msg,  []string) {
		return m,  addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
