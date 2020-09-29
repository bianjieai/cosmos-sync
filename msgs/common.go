package msgs

import (
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateMsgDocInfo(msg sdk.Msg, handler func() (Msg, []string)) MsgDocInfo {
	var (
		docTxMsg models.DocTxMsg
		signers  []string
		addrs    []string
	)

	m, addrcollections := handler()

	m.BuildMsg(msg)
	docTxMsg = models.DocTxMsg{
		Type: m.GetType(),
		Msg:  m,
	}

	_, signers = models.BuildDocSigners(msg.GetSigners())
	addrs = append(addrs, signers...)
	addrs = append(addrs, addrcollections...)

	return MsgDocInfo{
		DocTxMsg: docTxMsg,
		Signers:  signers,
		Addrs:    addrs,
	}
}

func ConvertMsg(v interface{}, msg interface{}) {
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
}
