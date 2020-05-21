package msgs

import (
	"github.com/bianjieai/irita-sync/models"
)

const (
	MsgTypeSend            = "send"
	MsgTypeNFTMint         = "nft_mint"
	MsgTypeNFTEdit         = "nft_edit"
	MsgTypeNFTTransfer     = "nft_transfer"
	MsgTypeNFTBurn         = "nft_burn"
	MsgTypeServiceDef      = "service_define"
	MsgTypeServiceBind     = "service_bind"
	MsgTypeServiceRequest  = "service_request"
	MsgTypeServiceResponse = "service_response"
	MsgTypeRecordCreate    = "create_record"
)

type (
	MsgDoc interface {
		GetType() string
		BuildMsg(v interface{})
		HandleTxMsg(msg interface{}) MsgDocInfo
	}

	MsgDocInfo struct {
		From       string        // parse from first msg
		To         string        // parse from first msg
		Coins      []models.Coin // parse from first msg
		Signer     string        // parse from first signer
		DocTxMsg   models.DocTxMsg
		ComplexMsg bool
		Signers    []string
	}
)
