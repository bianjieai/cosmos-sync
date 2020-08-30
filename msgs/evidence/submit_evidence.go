package evidence

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"encoding/json"
)

// MsgSubmitEvidence defines an sdk.Msg type that supports submitting arbitrary
// Evidence.
type DocMsgSubmitEvidence struct {
	Submitter string `bson:"submitter"`
	Evidence  string `bson:"evidence"`
}

func (m *DocMsgSubmitEvidence) GetType() string {
	return MsgTypeSubmitEvidence
}

func (m *DocMsgSubmitEvidence) BuildMsg(v interface{}) {
	msg := v.(MsgSubmitEvidence)
	m.Submitter = msg.Submitter.String()
	evidence, _ := json.Marshal(msg.Evidence)
	m.Evidence = string(evidence)

}

func (m *DocMsgSubmitEvidence) HandleTxMsg(msg MsgSubmitEvidence) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, msg.Submitter.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
