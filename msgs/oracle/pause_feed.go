package oracle

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgPauseFeed struct {
	FeedName string `bson:"feed_name" yaml:"feed_name"`
	Creator  string `bson:"creator"`
}

func (m *DocMsgPauseFeed) GetType() string {
	return TxTypePauseFeed
}

func (m *DocMsgPauseFeed) BuildMsg(v interface{}) {

	msg := v.(*MsgPauseFeed)

	m.FeedName = msg.FeedName
	m.Creator = msg.Creator
}

func (m *DocMsgPauseFeed) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgPauseFeed
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
