package ibc
//
//import (
//	"github.com/cosmos/cosmos-sdk/types"
//	. "github.com/bianjieai/irita-sync/msgs"
//)
//
//func HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
//	var (
//		msgDocInfo MsgDocInfo
//		//ibcClient  models.IbcClient
//	)
//	ok := true
//	//ibcClient.UpdateAt = timestamp
//	switch v.Type() {
//	case new(MsgRecvPacket).Type():
//		docMsg := DocMsgRecvPacket{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//	case new(MsgCreateClient).Type():
//		docMsg := DocMsgCreateClient{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//		//ibcClient.ConsensusState = docMsg.ConsensusState
//		//ibcClient.ClientState = docMsg.ClientState
//		//ibcClient.ClientId = docMsg.ClientID
//		//ibcClient.Signer = docMsg.Signer
//		//ibcClient.CreateAt = timestamp
//	case new(MsgUpdateClient).Type():
//		docMsg := DocMsgUpdateClient{}
//		msgDocInfo = docMsg.HandleTxMsg(v)
//		//ibcClient.Header = docMsg.Header
//		//ibcClient.ClientId = docMsg.ClientID
//		//ibcClient.Signer = docMsg.Signer
//	default:
//		ok = false
//	}
//	return msgDocInfo, ok
//}
