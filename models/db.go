package models

import (
	"context"
	"fmt"
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
)

var (
	_ctx  = context.Background()
	_conf *config.Config
	_cli  *qmgo.Client
)

func GetDbConf() *config.DataBaseConf {
	if _conf == nil {
		logger.Fatal("db.Init not work")
	}
	return &_conf.DataBase
}
func GetSrvConf() *config.ServerConf {
	if _conf == nil {
		logger.Fatal("db.Init not work")
	}
	return &_conf.Server
}

func GetClient() *qmgo.Client {
	return _cli
}

func Init(conf *config.Config) {
	_conf = conf
	var maxPoolSize uint64 = 4096
	// PrimaryMode indicates that only a primary is considered for reading. This is the default mode.
	client, err := qmgo.NewClient(_ctx, &qmgo.Config{
		Uri:         conf.DataBase.NodeUri,
		Database:    conf.DataBase.Database,
		MaxPoolSize: &maxPoolSize,
	})
	if err != nil {
		logger.Fatal(fmt.Sprintf("connect mongo failed, uri: %s, err:%s", conf.DataBase.NodeUri, err.Error()))
	}
	_cli = client

	logger.Info("init db success")

	// ensure table indexes
	ensureDocsIndexes()
	return
}

func Close() {
	logger.Info("release resource :mongoDb")
	if _cli != nil {
		_cli.Close(_ctx)
	}
}

func ensureIndexes(collectionName string, indexes []options.IndexModel) {
	c := _cli.Database(GetDbConf().Database).Collection(collectionName)
	if len(indexes) > 0 {
		for _, v := range indexes {
			if err := c.CreateOneIndex(context.Background(), v); err != nil {
				logger.Warn("ensure index fail", logger.String("collectionName", collectionName),
					logger.String("index", utils.MarshalJsonIgnoreErr(v)),
					logger.String("err", err.Error()))
			}
		}
	}
}

// get collection object
func ExecCollection(collectionName string, s func(*qmgo.Collection) error) error {
	c := _cli.Database(GetDbConf().Database).Collection(collectionName)
	return s(c)
}
