package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNameIbcClient = "ibc_client"

	IbcClientIdTag       = "client_id"
	IbcClientHeaderTag   = "header"
	IbcClientSignerTag   = "signer"
	IbcClientUpdateAtTag = "update_at"
)

type (
	IbcClient struct {
		ID             bson.ObjectId `bson:"_id"`
		ClientId       string        `bson:"client_id"`
		Header         Any           `bson:"header"`
		ClientState    Any           `bson:"client_state"`
		ConsensusState Any           `bson:"consensus_state"`
		Signer         string        `bson:"signer"`
		CreateAt       int64         `bson:"create_at"`
		UpdateAt       int64         `bson:"update_at"`
	}
	Any struct {
		// nolint
		TypeUrl string `bson:"type_url"`
		// Must be a valid serialized protocol buffer of the above specified type.
		Value string `bson:"value"`
	}
)

func (d IbcClient) Name() string {
	return CollectionNameIbcClient
}

func (d IbcClient) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{IbcClientIdTag},
		Unique:     true,
		Background: true,
	})

	ensureIndexes(d.Name(), indexes)
}

func (d IbcClient) PkKvPair() map[string]interface{} {
	return bson.M{IbcClientIdTag: d.ClientId}
}

func (m IbcClient) AllIbcClientMaps() (map[string]bson.ObjectId, error) {
	cond := bson.M{}
	var idens []IbcClient
	fn := func(c *mgo.Collection) error {
		return c.Find(cond).All(&idens)
	}
	if err := ExecCollection(m.Name(), fn); err != nil {
		return nil, err
	}
	mapData := make(map[string]bson.ObjectId, len(idens))
	for _, val := range idens {
		mapData[val.ClientId] = val.ID
	}
	return mapData, nil
}
