package record

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
)

type (
	DocMsgRecordCreate struct {
		Contents []Content `bson:"contents"`
		Creator  string    `bson:"creator"`
	}

	Content struct {
		Digest     string `bson:"digest"`
		DigestAlgo string `bson:"digest_algo"`
		URI        string `bson:"uri"`
		Meta       string `bson:"meta"`
	}
)

func (d *DocMsgRecordCreate) GetType() string {
	return MsgTypeRecordCreate
}

func (d *DocMsgRecordCreate) BuildMsg(msg interface{}) {
	m := msg.(*MsgRecordCreate)

	var docContents []Content
	if len(m.Contents) > 0 {
		for _, v := range m.Contents {
			docContents = append(docContents, Content{
				Digest:     v.Digest,
				DigestAlgo: v.DigestAlgo,
				URI:        v.URI,
				Meta:       v.Meta,
			})
		}
	}

	d.Contents = docContents
	d.Creator = m.Creator
}

func (m *DocMsgRecordCreate) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgRecordCreate)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return modules.CreateMsgDocInfo(v, handler)
}
