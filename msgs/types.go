package msgs

import (
	"github.com/bianjieai/irita-sync/models"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/irismod/nft"
	"github.com/irismod/service"
	"github.com/irismod/record"
	"github.com/irismod/token"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stake "github.com/cosmos/cosmos-sdk/x/staking/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	//coinswap "github.com/irismod/coinswap/types"
	//htlc "github.com/irismod/htlc/types"
)

const (
	MsgTypeSend         = "send"
	MsgTypeMultiSend    = "multisend"
	MsgTypeNFTMint      = "mint_nft"
	MsgTypeNFTEdit      = "edit_nft"
	MsgTypeNFTTransfer  = "transfer_nft"
	MsgTypeNFTBurn      = "burn_nft"
	MsgTypeIssueDenom   = "issue_denom"
	MsgTypeRecordCreate = "create_record"

	MsgTypeMintToken          = "mint_token"
	MsgTypeEditToken          = "edit_token"
	MsgTypeIssueToken         = "issue_token"
	MsgTypeTransferTokenOwner = "transfer_token_owner"

	MsgTypeDefineService             = "define_service"               // type for MsgDefineService
	MsgTypeBindService               = "bind_service"                 // type for MsgBindService
	MsgTypeUpdateServiceBinding      = "update_service_binding"       // type for MsgUpdateServiceBinding
	MsgTypeServiceSetWithdrawAddress = "service/set_withdraw_address" // type for MsgSetWithdrawAddress
	MsgTypeDisableServiceBinding     = "disable_service_binding"      // type for MsgDisableServiceBinding
	MsgTypeEnableServiceBinding      = "enable_service_binding"       // type for MsgEnableServiceBinding
	MsgTypeRefundServiceDeposit      = "refund_service_deposit"       // type for MsgRefundServiceDeposit
	MsgTypeCallService               = "call_service"                 // type for MsgCallService
	MsgTypeRespondService            = "respond_service"              // type for MsgRespondService
	MsgTypePauseRequestContext       = "pause_request_context"        // type for MsgPauseRequestContext
	MsgTypeStartRequestContext       = "start_request_context"        // type for MsgStartRequestContext
	MsgTypeKillRequestContext        = "kill_request_context"         // type for MsgKillRequestContext
	MsgTypeUpdateRequestContext      = "update_request_context"       // type for MsgUpdateRequestContext
	MsgTypeWithdrawEarnedFees        = "withdraw_earned_fees"         // type for MsgWithdrawEarnedFees

	MsgTypeStakeCreateValidator           = "create_validator"
	MsgTypeStakeEditValidator             = "edit_validator"
	MsgTypeStakeDelegate                  = "delegate"
	MsgTypeStakeBeginUnbonding            = "begin_unbonding"
	MsgTypeBeginRedelegate                = "begin_redelegate"
	MsgTypeUnjail                         = "unjail"
	MsgTypeSetWithdrawAddress             = "set_withdraw_address"
	MsgTypeWithdrawDelegatorReward        = "withdraw_delegator_reward"
	MsgTypeMsgFundCommunityPool           = "fund_community_pool"
	MsgTypeMsgWithdrawValidatorCommission = "withdraw_validator_commission"
	MsgTypeSubmitProposal                 = "submit_proposal"
	MsgTypeDeposit                        = "deposit"
	MsgTypeVote                           = "vote"

	TxTypeCreateHTLC = "CreateHTLC"
	TxTypeClaimHTLC  = "ClaimHTLC"
	TxTypeRefundHTLC = "RefundHTLC"

	MsgTypeCreateHTLC = "CreateHTLC"
	MsgTypeClaimHTLC  = "ClaimHTLC"
	MsgTypeRefundHTLC = "RefundHTLC"

	MsgTypeAddLiquidity    = "AddLiquidity"
	MsgTypeRemoveLiquidity = "RemoveLiquidity"
	MsgTypeSwapOrder       = "SwapOrder"

	MsgTypeSubmitEvidence  = "SubmitEvidence"
	MsgTypeVerifyInvariant = "VerifyInvariant"
)

type (
	MsgDocInfo struct {
		DocTxMsg models.DocTxMsg
		Addrs    []string
		Signers  []string
	}
	Msg models.Msg

	Coin models.Coin

	Coins []*Coin

	MsgSend = bank.MsgSend
	MsgMultiSend = bank.MsgMultiSend

	MsgNFTMint = nft.MsgMintNFT
	MsgNFTEdit = nft.MsgEditNFT
	MsgNFTTransfer = nft.MsgTransferNFT
	MsgNFTBurn = nft.MsgBurnNFT
	MsgIssueDenom = nft.MsgIssueDenom

	MsgDefineService = service.MsgDefineService
	MsgBindService = service.MsgBindService
	MsgCallService = service.MsgCallService
	MsgRespondService = service.MsgRespondService

	MsgUpdateServiceBinding = service.MsgUpdateServiceBinding
	MsgSetWithdrawAddress = service.MsgSetWithdrawAddress
	MsgDisableServiceBinding = service.MsgDisableServiceBinding
	MsgEnableServiceBinding = service.MsgEnableServiceBinding
	MsgRefundServiceDeposit = service.MsgRefundServiceDeposit
	MsgPauseRequestContext = service.MsgPauseRequestContext
	MsgStartRequestContext = service.MsgStartRequestContext
	MsgKillRequestContext = service.MsgKillRequestContext
	MsgUpdateRequestContext = service.MsgUpdateRequestContext
	MsgWithdrawEarnedFees = service.MsgWithdrawEarnedFees

	MsgRecordCreate = record.MsgCreateRecord

	MsgIssueToken = token.MsgIssueToken
	MsgEditToken = token.MsgEditToken
	MsgMintToken = token.MsgMintToken
	MsgTransferTokenOwner = token.MsgTransferTokenOwner

	MsgStakeCreate = stake.MsgCreateValidator
	MsgStakeEdit = stake.MsgEditValidator
	MsgStakeDelegate = stake.MsgDelegate
	MsgStakeBeginUnbonding = stake.MsgUndelegate
	MsgBeginRedelegate = stake.MsgBeginRedelegate
	MsgUnjail = slashing.MsgUnjail
	MsgStakeSetWithdrawAddress = dtypes.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward = distribution.MsgWithdrawDelegatorReward
	MsgFundCommunityPool = distribution.MsgFundCommunityPool
	MsgWithdrawValidatorCommission = distribution.MsgWithdrawValidatorCommission
	StakeValidator = stake.Validator
	Delegation = stake.Delegation
	UnbondingDelegation = stake.UnbondingDelegation

	MsgDeposit = gov.MsgDeposit
	MsgSubmitProposal = gov.MsgSubmitProposal
	TextProposal = gov.TextProposal
	MsgVote = gov.MsgVote
	Proposal = gov.Proposal
	SdkVote = gov.Vote

	//MsgSwapOrder = coinswap.MsgSwapOrder
	//MsgAddLiquidity = coinswap.MsgAddLiquidity
	//MsgRemoveLiquidity = coinswap.MsgRemoveLiquidity
	//
	//MsgClaimHTLC = htlc.MsgClaimHTLC
	//MsgCreateHTLC = htlc.MsgCreateHTLC
	//MsgRefundHTLC = htlc.MsgRefundHTLC

	MsgSubmitEvidence = evidence.MsgSubmitEvidence
	MsgVerifyInvariant = crisis.MsgVerifyInvariant

	MsgIBCTransfer = ibctransfer.MsgTransfer
	MsgIBCPacket = ibcchannel.MsgPacket
	MsgIBCTimeout = ibcchannel.MsgTimeout
	MsgUpdateClient = ibctendermint.MsgUpdateClient
	MsgCreateClient = ibctendermint.MsgCreateClient
	MsgSubmitClientMisbehaviour = ibctendermint.MsgSubmitClientMisbehaviour
)
