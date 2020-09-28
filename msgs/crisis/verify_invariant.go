package crisis
//
//import (
//	. "github.com/bianjieai/irita-sync/msgs"
//
//)
//
//type DocMsgVerifyInvariant struct {
//	Sender              string `bson:"sender"`
//	InvariantModuleName string `bson:"invariant_module_name" yaml:"invariant_module_name"`
//	InvariantRoute      string `bson:"invariant_route" yaml:"invariant_route"`
//}
//
//func (m *DocMsgVerifyInvariant) GetType() string {
//	return MsgTypeVerifyInvariant
//}
//
//func (m *DocMsgVerifyInvariant) BuildMsg(v interface{}) {
//	msg := v.(MsgVerifyInvariant)
//	m.Sender = msg.Sender.String()
//	m.InvariantModuleName= msg.InvariantModuleName
//	m.InvariantRoute = msg.InvariantRoute
//
//}
//
//func (m *DocMsgVerifyInvariant) HandleTxMsg(msg MsgVerifyInvariant) MsgDocInfo {
//
//	var (
//		addrs []string
//	)
//
//	addrs = append(addrs, msg.Sender.String())
//	handler := func() (Msg, []string) {
//		return m, addrs
//	}
//
//	return CreateMsgDocInfo(msg, handler)
//}
