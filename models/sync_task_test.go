package models

import (
	"encoding/json"
	"testing"
)

func TestSyncTask_GetExecutableTask(t *testing.T) {
	d := SyncTask{}

	if res, err := d.GetExecutableTask(120); err != nil {
		t.Fatal(err)
	} else {
		resBytes, _ := json.Marshal(res)
		t.Log(string(resBytes))
	}
}

func TestSyncTask_QueryAll(t *testing.T) {
	ret, err := new(SyncTask).QueryAll([]string{
		SyncTaskStatusUnHandled,
		SyncTaskStatusUnderway,
	},
		SyncTaskTypeCatchUp)

	if err != nil {
		t.Fatal(err)
	} else {
		resBytes, _ := json.Marshal(ret)
		t.Log(string(resBytes))
	}
}
