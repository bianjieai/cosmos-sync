package msgparser

const (
	BankRouteKey           string = "bank"
	StakingRouteKey        string = "staking"
	DistributionRouteKey   string = "distribution"
	CrisisRouteKey         string = "crisis"
	EvidenceRouteKey       string = "evidence"
	GovRouteKey            string = "gov"
	FeegrantRouteKey       string = "feegrant"
	CosmosSlashingRouteKey string = "cosmos_slashing"
	IritaSlashingRouteKey  string = "irita_slashing"
	NftRouteKey            string = "nft"
	MtRouteKey             string = "mt"
	ServiceRouteKey        string = "service"
	TokenRouteKey          string = "token"
	HtlcRouteKey           string = "htlc"
	CoinswapRouteKey       string = "coinswap"
	RandomRouteKey         string = "random"
	OracleRouteKey         string = "oracle"
	RecordRouteKey         string = "record"
	IdentityRouteKey       string = "identity"
	TIbcTransferRouteKey   string = "NFT"
	TIbcRouteKey           string = "tibc"
	EvmRouteKey            string = "evm"
	PermRouteKey           string = "perm"
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
