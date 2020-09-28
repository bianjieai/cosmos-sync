package record

type Content struct {
	DigestAlgo string `json:"digest_algo"`
	Digest     string `json:"digest"`
	URI        string `json:"uri"`
	Meta       string `json:"meta"`
}

type Packet struct {
	ID        string     `json:"id"`
	Timestamp uint64     `json:"timestamp,string"`
	Height    uint64     `json:"height,string"`
	TxHash    string     `json:"tx_hash"`
	Contents  []*Content `json:"contents"`
	Creator   string     `json:"creator"`
}

type IBCRecord struct {
	Type  string `json:"type"`
	Value Packet `json:"value"`
}


