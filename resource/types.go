package resource

type ChainRegisterResp struct {
	Schema       string `json:"$schema"`
	ChainName    string `json:"chain_name"`
	Status       string `json:"status"`
	NetworkType  string `json:"network_type"`
	PrettyName   string `json:"pretty_name"`
	ChainId      string `json:"chain_id"`
	Bech32Prefix string `json:"bech32_prefix"`
	DaemonName   string `json:"daemon_name"`
	NodeHome     string `json:"node_home"`
	Apis         struct {
		Rpc []struct {
			Address  string `json:"address"`
			Provider string `json:"provider"`
		} `json:"rpc"`
	} `json:"apis"`
}
