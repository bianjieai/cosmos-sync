package integration

import (
	"github.com/bianjieai/cosmos-sync/handlers"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/stretchr/testify/require"
)

func (s *IntegrationTestSuite) TestParseTxs() {
	block := int64(1201)
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()

	blockDoc, txDocs, err := handlers.ParseBlockAndTxs(block, c)
	require.NoError(s.T(), err)
	s.T().Log(utils.MarshalJsonIgnoreErr(blockDoc))
	s.T().Log(utils.MarshalJsonIgnoreErr(txDocs))
}
