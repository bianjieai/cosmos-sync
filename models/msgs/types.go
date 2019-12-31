package msgs

import (
	"gitlab.bianjie.ai/irita/ex-sync/models"
)

const (
	MsgTypeSend = "send"
)

type (
	MsgDocInfo struct {
		ComplexMsg bool
		From       string        // parse from first msg
		To         string        // parse from first msg
		Coins      []models.Coin // parse from first msg
		Signer     string        // parse from first signer
		Signers    []string
		DocTxMsg   models.DocTxMsg
	}

	CoinStr struct {
		Denom  string
		Amount string
	}
)
