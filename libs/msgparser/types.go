package msgparser

const (
	BankRouteKey         string = "bank"
	ServiceRouteKey      string = "service"
	NftRouteKey          string = "nft"
	MtRouteKey           string = "mt"
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
	FeegrantRouteKey     string = "feegrant"
	RandomRouteKey       string = "random"
	OracleRouteKey       string = "oracle"
	FarmRouteKey         string = "farm"
	IbcRouteKey          string = "ibc"
	IbcTransferRouteKey  string = "transfer"
	TIbcTransferRouteKey string = "NFT"
	TIbcRouteKey         string = "tibc"
	AuthzRouteKey        string = "authz"
	GroupRouteKey        string = "group"
)

type LegacyTx struct {
	// nonce corresponds to the account nonce (transaction sequence).
	Nonce uint64 `json:"nonce,omitempty"`
	// gas price defines the value for each gas unit
	GasPrice string `json:"gas_price,omitempty"`
	// gas defines the gas limit defined for the transaction.
	GasLimit uint64 `json:"gas,omitempty"`
	// hex formatted address of the recipient
	To string `json:"to,omitempty"`
	// value defines the unsigned integer value of the transaction amount.
	Amount string ` json:"value,omitempty"`
	// input defines the data payload bytes of the transaction.
	Data []byte `json:"data,omitempty"`
	// v defines the signature value
	V []byte `json:"v,omitempty"`
	// r defines the signature value
	R []byte `json:"r,omitempty"`
	// s define the signature value
	S []byte `json:"s,omitempty"`
}
