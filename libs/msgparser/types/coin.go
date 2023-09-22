package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Coin struct {
	Denom  string `bson:"denom" json:"denom"`
	Amount string `bson:"amount" json:"amount"`
}

type Coins []Coin

func BuildDocCoins(coins sdk.Coins) []Coin {
	var (
		res []Coin
	)
	if len(coins) > 0 {
		for _, v := range coins {
			c := Coin{
				Denom:  v.Denom,
				Amount: v.Amount.String(),
			}
			res = append(res, c)
		}
	}

	return res
}

func BuildDocCoin(coin sdk.Coin) Coin {
	return Coin{
		Denom:  coin.Denom,
		Amount: coin.Amount.String(),
	}
}

func BuildDocSigners(signers []sdk.AccAddress) (string, []string) {
	var (
		firstSigner string
		allSigners  []string
	)
	if len(signers) == 0 {
		return firstSigner, allSigners
	}
	for _, v := range signers {
		if firstSigner == "" {
			firstSigner = v.String()
		}

		allSigners = append(allSigners, v.String())
	}

	return firstSigner, allSigners
}
