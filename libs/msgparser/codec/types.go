package codec

import (
	okexchain "github.com/okex/exchain/app/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
)

func SetBech32Prefix() {
	config := sdk.GetConfig()
	okexchain.SetBech32Prefixes(config)
	config.Seal()
}
