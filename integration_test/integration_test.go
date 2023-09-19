package integration

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/handlers"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type IntegrationTestSuite struct {
	*config.Config
	suite.Suite
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

func TestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	os.Setenv(constant.EnvNameConfigFilePath, "../config/config.toml")

	config.InitEnv()
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	models.Init(cfg)
	pool.Init(cfg)
	handlers.InitMsgParser()
	s.Config = cfg
}

func (s *IntegrationTestSuite) TearDownSuite() {
	models.Close()
}
