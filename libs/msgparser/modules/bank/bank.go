package bank

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type BankClient struct {
}

func NewClient() Client {
	return BankClient{}
}

func (bank BankClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgSend:
		docMsg := DocMsgSend{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgMultiSend:
		docMsg := DocMsgMultiSend{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
