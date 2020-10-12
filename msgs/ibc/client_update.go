package ibc
//
//import (
//	. "github.com/bianjieai/irita-sync/msgs"
//	"github.com/bianjieai/irita-sync/utils"
//	"github.com/bianjieai/irita-sync/models"
//)
//
//// MsgUpdateClient defines a message to update an IBC client
//type DocMsgUpdateClient struct {
//	ClientID string     `bson:"client_id" yaml:"client_id"`
//	Header   models.Any `bson:"header" yaml:"header"`
//	Signer   string     `bson:"signer" yaml:"signer"`
//}
//
//func (m *DocMsgUpdateClient) GetType() string {
//	return MsgTypeUpdateClient
//}
//
//func (m *DocMsgUpdateClient) BuildMsg(v interface{}) {
//	msg := v.(MsgUpdateClient)
//
//	m.ClientID = msg.ClientID
//	m.Signer = msg.Signer.String()
//	m.Header = models.Any{TypeUrl: msg.Header.GetTypeUrl(), Value: string(msg.Header.GetValue())}
//}
//
//func (m *DocMsgUpdateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
//	var (
//		addrs []string
//		msg   MsgUpdateClient
//	)
//
//	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
//	addrs = append(addrs, msg.Signer.String())
//	handler := func() (Msg, []string) {
//		return m, addrs
//	}
//
//	return CreateMsgDocInfo(v, handler)
//}
