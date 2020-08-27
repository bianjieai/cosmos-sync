package evidence

import (
	. "github.com/bianjieai/irita-sync/msgs"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"encoding/json"
)

// MsgSubmitEvidence defines an sdk.Msg type that supports submitting arbitrary
// Evidence.
type DocMsgSubmitEvidence struct {
	Submitter string `bson:"submitter"`
	Evidence  string `bson:"evidence"`
}

func (m *DocMsgSubmitEvidence) GetType() string {
	return TxTypeSubmitEvidence
}

func (m *DocMsgSubmitEvidence) BuildMsg(v interface{}) {
	msg := v.(MsgSubmitEvidence)
	m.Submitter = msg.Submitter.String()
	evidence, _ := json.Marshal(msg.Evidence)
	m.Evidence = string(evidence)

}

func (m *DocMsgSubmitEvidence) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.Submitter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
