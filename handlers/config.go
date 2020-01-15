package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	conf "github.com/bianjieai/irita-sync/confs/server"
	iconfig "github.com/bianjieai/irita/config"
)

func init() {
	config := sdk.GetConfig()
	iconfig.SetNetworkType(conf.SvrConf.NetWork)
	iritaConfig := iconfig.GetConfig()
	config.SetBech32PrefixForAccount(iritaConfig.GetBech32AccountAddrPrefix(), iritaConfig.GetBech32AccountPubPrefix())
	config.SetBech32PrefixForValidator(iritaConfig.GetBech32ValidatorAddrPrefix(), iritaConfig.GetBech32ValidatorPubPrefix())
	config.SetBech32PrefixForConsensusNode(iritaConfig.GetBech32ConsensusAddrPrefix(), iritaConfig.GetBech32ConsensusPubPrefix())
	config.Seal()
}
