package models

import (
	"fmt"
	"gitlab.bianjie.ai/irita/ex-sync/confs/db"
	"gitlab.bianjie.ai/irita/ex-sync/libs/logger"
	"gitlab.bianjie.ai/irita/ex-sync/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"strings"
	"time"
)

var (
	session *mgo.Session
)

func init() {
	addrs := strings.Split(db.DbConf.Addrs, ",")
	dialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Database:  db.DbConf.Database,
		Username:  db.DbConf.User,
		Password:  db.DbConf.Passwd,
		Direct:    true,
		Timeout:   time.Second * 10,
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	var err error
	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		logger.Fatal("connect db fail", logger.String("err", err.Error()))
	}
	session.SetMode(mgo.Strong, true)
	logger.Info("init db success")

	// ensure table indexes
	ensureDocsIndexes()
}

func Close() {
	logger.Info("release resource :mongoDb")
	session.Close()
}

func getSession() *mgo.Session {
	// max session num is 4096
	return session.Clone()
}

func ensureIndexes(collectionName string, indexes []mgo.Index) {
	c := session.DB(db.DbConf.Database).C(collectionName)
	if len(indexes) > 0 {
		for _, v := range indexes {
			if err := c.EnsureIndex(v); err != nil {
				logger.Warn("ensure index fail", logger.String("collectionName", collectionName),
					logger.String("index", utils.MarshalJsonIgnoreErr(v)),
					logger.String("err", err.Error()))
			}
		}
	}
}

// get collection object
func ExecCollection(collectionName string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(db.DbConf.Database).C(collectionName)
	return s(c)
}

func Save(h Docs) error {
	save := func(c *mgo.Collection) error {
		pk := h.PkKvPair()
		n, _ := c.Find(pk).Count()
		if n >= 1 {
			return fmt.Errorf("record exist")
		}
		return c.Insert(h)
	}
	return ExecCollection(h.Name(), save)
}

func Update(h Docs) error {
	update := func(c *mgo.Collection) error {
		key := h.PkKvPair()
		return c.Update(key, h)
	}
	return ExecCollection(h.Name(), update)
}

func Delete(h Docs) error {
	remove := func(c *mgo.Collection) error {
		key := h.PkKvPair()
		return c.Remove(key)
	}
	return ExecCollection(h.Name(), remove)
}

//mgo transaction method
//detail to see: https://godoc.org/gopkg.in/mgo.v2/txn
func Txn(ops []txn.Op) error {
	session := getSession()
	defer session.Close()

	c := session.DB(db.DbConf.Database).C(CollectionNameTxn)
	runner := txn.NewRunner(c)

	txObjectId := bson.NewObjectId()
	err := runner.Run(ops, txObjectId, nil)
	if err != nil {
		if err == txn.ErrAborted {
			err = runner.Resume(txObjectId)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
