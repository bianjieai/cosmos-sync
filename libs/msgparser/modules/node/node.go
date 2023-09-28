package node

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type NodeClient struct {
}

func NewClient() NodeClient {
	return NodeClient{}
}

func (nft NodeClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgNodeCreate:
		docMsg := DocMsgCreateValidator{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgNodeUpdate:
		docMsg := DocMsgUpdateValidator{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgNodeRemove:
		docMsg := DocMsgRemoveValidator{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgNodeGrant:
		docMsg := DocMsgGrantNode{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgNodeRevoke:
		docMsg := DocMsgRevokeNode{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
