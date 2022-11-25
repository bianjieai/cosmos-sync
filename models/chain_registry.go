package models

import (
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type ChainRegistry struct {
	Chain        string `bson:"chain"`
	ChainJsonUrl string `bson:"chain_json_url"`
}

func (c ChainRegistry) Name() string {
	return "chain_registry"
}

func (c ChainRegistry) PkKvPair() map[string]interface{} {
	return bson.M{"chain": c.Chain}
}
func (c ChainRegistry) EnsureIndexes() {

}

func (c ChainRegistry) FindOne(chain string) (ChainRegistry, error) {
	var res ChainRegistry
	fn := func(c *qmgo.Collection) error {
		return c.Find(_ctx, bson.M{
			"chain": chain,
		}).One(&res)
	}

	err := ExecCollection(c.Name(), fn)
	return res, err
}
