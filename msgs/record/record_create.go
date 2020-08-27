package record

import (
	. "github.com/bianjieai/irita-sync/msgs"
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
	m := msg.(MsgRecordCreate)

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
	d.Creator = m.Creator.String()
}

func (m *DocMsgRecordCreate) HandleTxMsg(msg MsgRecordCreate) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Creator)
	handler := func() (Msg,  []string) {
		return m,  addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
