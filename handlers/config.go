package handlers

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
)

func initBech32Prefix(bech32AccPrefix string) {
	if bech32AccPrefix != "" {
		codec.SetBech32Prefixs(bech32AccPrefix)
	}
}
