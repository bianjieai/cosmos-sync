package codec

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = "iaa"
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = "iap"
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = "iva"
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = "ivp"
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = "ica"
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = "icp"
)

func SetBech32Prefix(bech32PrefixAccAddr, bech32PrefixAccPub, bech32PrefixValAddr,
	bech32PrefixValPub, bech32PrefixConsAddr, bech32PrefixConsPub string) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(bech32PrefixAccAddr, bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(bech32PrefixValAddr, bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(bech32PrefixConsAddr, bech32PrefixConsPub)
	config.Seal()
}

func SetBech32Prefixs(prefixs []string) {
	var (
		bech32PrefixAccAddr []string
		bech32PrefixAccPub  []string

		bech32PrefixValAddr []string
		bech32PrefixValPub  []string

		bech32PrefixConsAddr []string
		bech32PrefixConsPub  []string
	)
	const (
		PrefixValidator = "val"
		PrefixConsensus = "cons"
		PrefixPublic    = "pub"
		PrefixOperator  = "oper"
	)
	for _, val := range prefixs {
		Bech32PrefixAccAddr = val
		Bech32PrefixAccPub = Bech32PrefixAccAddr + PrefixPublic
		Bech32PrefixValAddr = Bech32PrefixAccAddr + PrefixValidator + PrefixOperator
		Bech32PrefixValPub = Bech32PrefixAccAddr + PrefixValidator + PrefixOperator + PrefixPublic
		Bech32PrefixConsAddr = Bech32PrefixAccAddr + PrefixValidator + PrefixConsensus
		Bech32PrefixConsPub = Bech32PrefixAccAddr + PrefixValidator + PrefixConsensus + PrefixPublic

		bech32PrefixAccAddr = append(bech32PrefixAccAddr, Bech32PrefixAccAddr)
		bech32PrefixAccPub = append(bech32PrefixAccPub, Bech32PrefixAccPub)
		bech32PrefixValAddr = append(bech32PrefixValAddr, Bech32PrefixValAddr)
		bech32PrefixValPub = append(bech32PrefixValPub, Bech32PrefixValPub)
		bech32PrefixConsAddr = append(bech32PrefixConsAddr, Bech32PrefixConsAddr)
		bech32PrefixConsPub = append(bech32PrefixConsPub, Bech32PrefixConsPub)
	}

	config := sdk.GetConfig()
	config.SetBech32PrefixesForAccount(bech32PrefixAccAddr, bech32PrefixAccPub)
	config.SetBech32PrefixesForValidator(bech32PrefixValAddr, bech32PrefixValPub)
	config.SetBech32PrefixesForConsensusNode(bech32PrefixConsAddr, bech32PrefixConsPub)
	config.Seal()
}
