package wasm

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgStoreCode struct {
	// Sender is the that actor that signed the messages
	Sender string `bson:"sender"`
	// WASMByteCode can be raw or gzip compressed
	WASMByteCode []byte `bson:"wasm_byte_code"`
	// Source is a valid absolute HTTPS URI to the contract's source code, optional
	Source string `bson:"source"`
	// Builder is a valid docker image name with tag, optional
	Builder string `bson:"builder"`
	// InstantiatePermission access control to apply on contract creation, optional
	InstantiatePermission AccessConfig `bson:"instantiate_permission"`
}

type AccessConfig struct {
	Permission string `bson:"permission"`
	Address    string `bson:"address"`
}

func (m *DocMsgStoreCode) GetType() string {
	return MsgTypeStoreCode
}

func (m *DocMsgStoreCode) BuildMsg(v interface{}) {
	msg := v.(*MsgStoreCode)
	m.Sender = msg.Sender
	m.WASMByteCode = msg.WASMByteCode
	m.Source = msg.Source
	m.Builder = msg.Builder
	if msg.InstantiatePermission != nil {
		m.InstantiatePermission = AccessConfig{
			Permission: msg.InstantiatePermission.Permission.String(),
			Address: msg.InstantiatePermission.Address,
		}
	}

}

func (m *DocMsgStoreCode) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgStoreCode
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
