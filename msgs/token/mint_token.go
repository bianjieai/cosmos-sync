package token

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgMintToken struct {
	Symbol string `bson:"symbol"`
	Amount uint64 `bson:"amount"`
	To     string `bson:"to"`
	Owner  string `bson:"owner"`
}

func (m *DocMsgMintToken) GetType() string {
	return MsgTypeMintToken
}

func (m *DocMsgMintToken) BuildMsg(v interface{}) {
	msg := v.(*MsgMintToken)

	m.Symbol = msg.Symbol
	m.Amount = msg.Amount
	m.To = msg.To.String()
	m.Owner = msg.Owner.String()
}

func (m *DocMsgMintToken) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgMintToken
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Owner.String(), msg.To.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
