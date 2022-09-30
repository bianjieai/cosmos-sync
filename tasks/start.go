package tasks

import "github.com/bianjieai/cosmos-sync/monitor"

func Start(synctask SyncTask) {
	go synctask.StartCreateTask()
	go synctask.StartExecuteTask()
	go monitor.Start()
}
