package msgparser

const (
	BankRouteKey         string = "bank"
	ServiceRouteKey      string = "service"
	NftRouteKey          string = "nft"
	RecordRouteKey       string = "record"
	TokenRouteKey        string = "token"
	CoinswapRouteKey     string = "coinswap"
	CrisisRouteKey       string = "crisis"
	DistributionRouteKey string = "distribution"
	SlashingRouteKey     string = "slashing"
	EvidenceRouteKey     string = "evidence"
	HtlcRouteKey         string = "htlc"
	StakingRouteKey      string = "staking"
	GovRouteKey          string = "gov"
	RandomRouteKey       string = "random"
	OracleRouteKey       string = "oracle"
	FarmRouteKey         string = "farm"
	IbcRouteKey          string = "ibc"
	IbcTransferRouteKey  string = "transfer"
	TIbcTransferRouteKey string = "nftTransfer"
	TIbcRouteKey         string = "tibc"
)

var RouteHandlerMap = map[string]Handler{
	BankRouteKey:         handleBank,
	ServiceRouteKey:      handleService,
	NftRouteKey:          handleNft,
	RecordRouteKey:       handleRecord,
	TokenRouteKey:        handleToken,
	CoinswapRouteKey:     handleCoinswap,
	CrisisRouteKey:       handleCrisis,
	DistributionRouteKey: handleDistribution,
	SlashingRouteKey:     handleSlashing,
	EvidenceRouteKey:     handleEvidence,
	HtlcRouteKey:         handleHtlc,
	StakingRouteKey:      handleStaking,
	GovRouteKey:          handleGov,
	RandomRouteKey:       handleRandom,
	OracleRouteKey:       handleOracle,
	IbcRouteKey:          handleIbc,
	IbcTransferRouteKey:  handleIbc,
	FarmRouteKey:         handleFarm,
	TIbcTransferRouteKey: handleTIbc,
	TIbcRouteKey:         handleTIbc,
}
