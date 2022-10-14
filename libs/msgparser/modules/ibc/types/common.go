package types

import (
	codec "github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	models "github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
	"github.com/okex/exchain/libs/cosmos-sdk/codec/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
)

func CreateMsgDocInfo(msg sdk.Msg, handler func() (Msg, []string)) MsgDocInfo {
	var (
		docTxMsg models.TxMsg
		signers  []string
		addrs    []string
	)

	m, addrcollections := handler()

	m.BuildMsg(msg)
	docTxMsg = models.TxMsg{
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

// ConvertMsg
// Deprecated, use type assertion instead of this
func ConvertMsg(v interface{}, msg interface{}) {
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
}

func ConvertAny(v *types.Any) string {
	if v != nil {
		return utils.MarshalJsonIgnoreErr(v)
	}
	return ""
}

func UnmarshalAcknowledgement(bytesdata []byte) string {
	var result Acknowledgement
	codec.GetCodec().MustUnmarshalJSON(bytesdata, &result)
	return result.String()
}
