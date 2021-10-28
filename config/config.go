package config

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	"github.com/spf13/viper"
	"os"
)

var (
	ConfigFilePath string
)

type (
	Config struct {
		DataBase DataBaseConf `mapstructure:"database"`
		Server   ServerConf   `mapstructure:"server"`
	}
	DataBaseConf struct {
		Addrs    string `mapstructure:"addrs"`
		User     string `mapstructure:"user"`
		Passwd   string `mapstructure:"passwd" json:"-"`
		Database string `mapstructure:"database"`
	}

	ServerConf struct {
		NodeUrls                  string `mapstructure:"node_urls"`
		WorkerNumCreateTask       int    `mapstructure:"worker_num_create_task"`
		WorkerNumExecuteTask      int    `mapstructure:"worker_num_execute_task"`
		WorkerMaxSleepTime        int    `mapstructure:"worker_max_sleep_time"`
		BlockNumPerWorkerHandle   int    `mapstructure:"block_num_per_worker_handle"`
		SleepTimeCreateTaskWorker int    `mapstructure:"sleep_time_create_task_worker"`

		MaxConnectionNum   int    `mapstructure:"max_connection_num"`
		InitConnectionNum  int    `mapstructure:"init_connection_num"`
		Bech32AccPrefix    string `mapstructure:"bech32_acc_prefix"`
		ChainId            string `mapstructure:"chain_id"`
		ChainBlockInterval int    `mapstructure:"chain_block_interval"`
		BehindBlockNum     int    `mapstructure:"behind_block_num"`

		PromethousPort string `mapstructure:"promethous_port"`
		SupportModules string `mapstructure:"support_modules"`
		DenyModules    string `mapstructure:"deny_modules"`
	}
)

func init() {
	websit, found := os.LookupEnv(constant.EnvNameConfigFilePath)
	if found {
		ConfigFilePath = websit
	} else {
		panic("not found CONFIG_FILE_PATH")
	}
}

func ReadConfig() (*Config, error) {

	rootViper := viper.New()
	// Find home directory.
	rootViper.SetConfigFile(ConfigFilePath)

	// Find and read the config file
	if err := rootViper.ReadInConfig(); err != nil { // Handle errors reading the config file
		return nil, err
	}

	var cfg Config
	if err := rootViper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// calculate sleep time of create task goroutine, time unit is second
	// 1. sleepTime must less than blockNumPerWorkerHandle*chainTimeInterval,
	//    otherwise create task goroutine can't succeed create follow task
	// 2. use value of sleepTime/5 to make create task worker do task in time
	sleepTimeCreateTaskWorker := (cfg.Server.BlockNumPerWorkerHandle * cfg.Server.ChainBlockInterval) / 5
	if sleepTimeCreateTaskWorker == 0 {
		sleepTimeCreateTaskWorker = 1
	}
	cfg.Server.SleepTimeCreateTaskWorker = sleepTimeCreateTaskWorker

	logger.Debug("config: " + utils.MarshalJsonIgnoreErr(cfg))

	return &cfg, nil
}
