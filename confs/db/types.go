package db

import (
	constant "gitlab.bianjie.ai/irita/ex-sync/confs"
	"gitlab.bianjie.ai/irita/ex-sync/libs/logger"
	"gitlab.bianjie.ai/irita/ex-sync/utils"
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
	addrs    = "192.168.150.31:27017"
	user     = "iris"
	passwd   = "irispassword"
	database = "rainbow-server"
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
