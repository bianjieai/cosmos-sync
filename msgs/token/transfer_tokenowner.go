package token

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgTransferTokenOwner struct {
	SrcOwner string `bson:"src_owner"`
	DstOwner string `bson:"dst_owner"`
	Symbol   string `bson:"symbol"`
}

func (m *DocMsgTransferTokenOwner) GetType() string {
	return MsgTypeTransferTokenOwner
}

func (m *DocMsgTransferTokenOwner) BuildMsg(v interface{}) {
	msg := v.(*MsgTransferTokenOwner)

	m.Symbol = msg.Symbol
	m.SrcOwner = msg.SrcOwner.String()
	m.DstOwner = msg.DstOwner.String()
}

func (m *DocMsgTransferTokenOwner) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTransferTokenOwner
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.SrcOwner.String(), msg.DstOwner.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
