package handlers

import "github.com/kaifei-bianjie/msg-parser/codec"

var (
	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr string
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub string
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr string
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub string
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr string
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub string
)

func initIrisPrefix() {
	const (
		Bech32ChainPrefix = "i"
		PrefixAcc         = "a"
		PrefixValidator   = "v"
		PrefixConsensus   = "c"
		PrefixPublic      = "p"
		PrefixAddress     = "a"
	)
	bech32AccPrefix := Bech32ChainPrefix
	Bech32PrefixAccAddr = bech32AccPrefix + PrefixAcc + PrefixAddress
	Bech32PrefixAccPub = bech32AccPrefix + PrefixAcc + PrefixPublic
	Bech32PrefixValAddr = bech32AccPrefix + PrefixValidator + PrefixAddress
	Bech32PrefixValPub = bech32AccPrefix + PrefixValidator + PrefixPublic
	Bech32PrefixConsAddr = bech32AccPrefix + PrefixConsensus + PrefixAddress
	Bech32PrefixConsPub = bech32AccPrefix + PrefixConsensus + PrefixPublic

	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
}

func initCosmosPrefix(bech32AccPrefix string) {
	const (
		PrefixValidator = "val"
		PrefixConsensus = "cons"
		PrefixPublic    = "pub"
		PrefixOperator  = "oper"
	)
	Bech32PrefixAccAddr = bech32AccPrefix
	Bech32PrefixAccPub = Bech32PrefixAccAddr + PrefixPublic
	Bech32PrefixValAddr = Bech32PrefixAccAddr + PrefixValidator + PrefixOperator
	Bech32PrefixValPub = Bech32PrefixAccAddr + PrefixValidator + PrefixOperator + PrefixPublic
	Bech32PrefixConsAddr = Bech32PrefixAccAddr + PrefixValidator + PrefixConsensus
	Bech32PrefixConsPub = Bech32PrefixAccAddr + PrefixValidator + PrefixConsensus + PrefixPublic

	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
}

func initBech32Prefix(bech32AccPrefix string) {
	if bech32AccPrefix != "" {
		if bech32AccPrefix == "iaa" {
			initIrisPrefix()
		} else {
			initCosmosPrefix(bech32AccPrefix)
		}
	}
}
