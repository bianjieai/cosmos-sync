package db

import (
	constant "github.com/bianjieai/irita-sync/confs"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/utils"
	"os"
)

type DBConf struct {
	Addrs    string
	User     string `json:"_"`
	Passwd   string `json:"_"`
	Database string
}

var (
	DbConf   *DBConf
	addrs    = "10.1.4.130:27217"
	user     = "iris"
	passwd   = "irispassword"
	database = "sync-iris"
)

// get value of env var
func init() {
	if v, ok := os.LookupEnv(constant.EnvNameDbAddr); ok {
		addrs = v
	}

	if v, ok := os.LookupEnv(constant.EnvNameDbUser); ok {
		user = v
	}

	if v, ok := os.LookupEnv(constant.EnvNameDbPassWd); ok {
		passwd = v
	}

	if v, ok := os.LookupEnv(constant.EnvNameDbDataBase); ok {
		database = v
	}

	DbConf = &DBConf{
		Addrs:    addrs,
		User:     user,
		Passwd:   passwd,
		Database: database,
	}

	logger.Debug("print db config",
		logger.String("dbConfig", utils.MarshalJsonIgnoreErr(DbConf)))
}
