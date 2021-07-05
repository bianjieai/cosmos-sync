package msgparser

const (
	BankRouteKey         string = "bank"
	StakingRouteKey      string = "staking"
	DistributionRouteKey string = "distribution"
	CrisisRouteKey       string = "crisis"
	EvidenceRouteKey     string = "evidence"
	GovRouteKey          string = "gov"
	SlashingRouteKey     string = "slashing"
	IbcRouteKey          string = "ibc"
	IbcTransferRouteKey  string = "transfer"
)

var RouteHandlerMap = map[string]Handler{
	BankRouteKey:         handleBank,
	StakingRouteKey:      handleStaking,
	DistributionRouteKey: handleDistribution,
	CrisisRouteKey:       handleCrisis,
	EvidenceRouteKey:     handleEvidence,
	GovRouteKey:          handleGov,
	SlashingRouteKey:     handleSlashing,
	IbcRouteKey:          handleIbc,
	IbcTransferRouteKey:  handleIbc,
}
