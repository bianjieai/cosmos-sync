package integration

import (
	"encoding/json"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/stretchr/testify/require"
)

func (s *IntegrationTestSuite) TestSyncTask_GetExecutableTask() {
	d := models.SyncTask{}

	res, err := d.GetExecutableTask(120)
	require.NoError(s.T(), err)
	resBytes, _ := json.Marshal(res)
	s.T().Log(string(resBytes))
}

func (s *IntegrationTestSuite) TestSyncTask_QueryAll() {
	ret, err := new(models.SyncTask).QueryAll([]string{
		models.SyncTaskStatusUnHandled,
		models.SyncTaskStatusUnderway,
	},
		models.SyncTaskTypeCatchUp)

	require.NoError(s.T(), err)
	resBytes, _ := json.Marshal(ret)
	s.T().Log(string(resBytes))
}
