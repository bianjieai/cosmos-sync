package oracle

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgStartFeed struct {
	FeedName string `bson:"feed_name" yaml:"feed_name"`
	Creator  string `bson:"creator"`
}

func (m *DocMsgStartFeed) GetType() string {
	return TxTypeStartFeed
}

func (m *DocMsgStartFeed) BuildMsg(v interface{}) {
	msg := v.(*MsgStartFeed)

	m.FeedName = msg.FeedName
	m.Creator = msg.Creator
}

func (m *DocMsgStartFeed) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgStartFeed
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
