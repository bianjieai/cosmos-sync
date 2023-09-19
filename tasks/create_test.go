package tasks

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/handlers"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	"os"
	"testing"
	"time"
)

var testCfg *config.Config

func TestMain(m *testing.M) {
	os.Setenv(constant.EnvNameConfigFilePath, "../config/config.toml")

	config.InitEnv()
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	models.Init(cfg)
	pool.Init(cfg)
	handlers.InitMsgParser()
	testCfg = cfg
	m.Run()
}

func TestSyncTaskService_createTask(t *testing.T) {
	s := syncTaskService{conf: testCfg}
	chanLimit := make(chan bool, 1)

	for {
		chanLimit <- true
		go s.createTask(100, chanLimit)
		time.Sleep(time.Duration(1) * time.Minute)
	}
}
