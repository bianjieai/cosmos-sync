package ibc

import (
	"encoding/json"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"gitlab.bianjie.ai/cschain/cschain/modules/ibc/core/types"
)

// MsgUpdateClient defines a message to update an IBC client
type DocMsgUpdateClient struct {
	ClientID string `bson:"client_id" yaml:"client_id"`
	Header   string `bson:"header" yaml:"header"`
	Signer   string `bson:"signer" yaml:"signer"`
}

func (m *DocMsgUpdateClient) GetType() string {
	return MsgTypeUpdateClient
}

func (m *DocMsgUpdateClient) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateClient)

	if header, err := types.UnpackHeader(msg.Header); err == nil {
		data, _ := json.Marshal(header)
		m.Header = string(data)
	}

	m.ClientID = msg.ClientID
	m.Signer = msg.Signer
}

func (m *DocMsgUpdateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUpdateClient)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return modules.CreateMsgDocInfo(v, handler)
}
