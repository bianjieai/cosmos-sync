// interface for a document

package models

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CollectionNameTxn = "sync_txn"
)

var (
	SyncTaskModel SyncTask
	BlockModel    Block
	TxModel       Tx

	Collections = []Docs{
		SyncTaskModel,
		BlockModel,
		TxModel,
	}
)

type (
	Docs interface {
		// collection name
		Name() string
		// ensure indexes
		EnsureIndexes()
		// primary key pair(used to find a unique record)
		PkKvPair() map[string]interface{}
	}

	Coin struct {
		Denom  string `bson:"denom"`
		Amount string `bson:"amount"`
	}

	Coins []*Coin
)

// Description
type Description struct {
	Moniker         string `bson:"moniker"`
	Identity        string `bson:"identity"`
	Website         string `bson:"website"`
	SecurityContact string `bson:"security_contact"`
	Details         string `bson:"details"`
}

type CommissionRates struct {
	Rate          string `bson:"rate"`            // the commission rate charged to delegators
	MaxRate       string `bson:"max_rate"`        // maximum commission rate which validator can ever charge
	MaxChangeRate string `bson:"max_change_rate"` // maximum daily increase of the validator commission
}

func ensureDocsIndexes() {
	if len(Collections) > 0 {
		for _, v := range Collections {
			v.EnsureIndexes()
		}
	}
}

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

func BuildFee(fee sdk.Coins, gas uint64) Fee {
	return Fee{
		Amount: BuildDocCoins(fee),
		Gas:    int64(gas),
	}
}
