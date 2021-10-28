package handlers

import (
	"github.com/kaifei-bianjie/msg-parser/codec"
)

const (
	// Bech32ChainPrefix defines the prefix of this chain
	Bech32ChainPrefix = "i"

	// PrefixAcc is the prefix for account
	PrefixAcc = "a"

	// PrefixValidator is the prefix for validator keys
	PrefixValidator = "v"

	// PrefixConsensus is the prefix for consensus keys
	PrefixConsensus = "c"

	// PrefixPublic is the prefix for public
	PrefixPublic = "p"

	// PrefixAddress is the prefix for address
	PrefixAddress = "a"
)

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

func initBech32Prefix(bech32AccPrefix string) {
	Bech32PrefixAccAddr = bech32AccPrefix + PrefixAcc + PrefixAddress
	Bech32PrefixAccPub = bech32AccPrefix + PrefixAcc + PrefixPublic
	Bech32PrefixValAddr = bech32AccPrefix + PrefixValidator + PrefixAddress
	Bech32PrefixValPub = bech32AccPrefix + PrefixValidator + PrefixPublic
	Bech32PrefixConsAddr = bech32AccPrefix + PrefixConsensus + PrefixAddress
	Bech32PrefixConsPub = bech32AccPrefix + PrefixConsensus + PrefixPublic

	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
}
