package codec

import sdk "github.com/cosmos/cosmos-sdk/types"

func SetBech32Prefix(bech32PrefixAccAddr, bech32PrefixAccPub, bech32PrefixValAddr,
	bech32PrefixValPub, bech32PrefixConsAddr, bech32PrefixConsPub string) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(bech32PrefixAccAddr, bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(bech32PrefixValAddr, bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(bech32PrefixConsAddr, bech32PrefixConsPub)
	config.Seal()
}
