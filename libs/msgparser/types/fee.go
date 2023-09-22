package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Fee struct {
	Amount []Coin `bson:"amount"`
	Gas    int64  `bson:"gas"`
}

func BuildFee(fee sdk.Coins, gas uint64) *Fee {
	return &Fee{
		Amount: BuildDocCoins(fee),
		Gas:    int64(gas),
	}
}
