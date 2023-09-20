package ibc

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	recordtype "gitlab.bianjie.ai/cschain/cschain/modules/ibc/applications/record/types"
	"gitlab.bianjie.ai/cschain/cschain/modules/ibc/core/types"
)

type Content struct {
	DigestAlgo string `bson:"digest_algo"`
	Digest     string `bson:"digest"`
	URI        string `bson:"uri"`
	Meta       string `bson:"meta"`
}

type PacketData struct {
	ID        string     `bson:"id"`
	Timestamp uint64     `bson:"timestamp"`
	Height    uint64     `bson:"height"`
	TxHash    string     `bson:"tx_hash"`
	Contents  []*Content `bson:"contents"`
	Creator   string     `bson:"creator"`
}

func DecodeToIBCRecord(packet types.Packet) (ibcRecord PacketData) {
	var value recordtype.Packet
	codec.GetMarshaler().UnmarshalJSON([]byte(packet.Data), &value)
	ibcRecord = PacketData{
		ID:       value.ID,
		Height:   value.Height,
		Creator:  value.Creator,
		TxHash:   value.TxHash,
		Contents: loadPacketContents(value.Contents),
	}
	return
}
func loadPacketContents(contents []*recordtype.Content) []*Content {
	sliceContent := make([]*Content, 0, len(contents))
	for _, val := range contents {
		sliceContent = append(sliceContent, &Content{
			Digest:     val.Digest,
			DigestAlgo: val.DigestAlgo,
			Meta:       val.Meta,
			URI:        val.URI,
		})
	}
	return sliceContent
}
