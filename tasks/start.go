package tasks

import "github.com/bianjieai/irita-sync/monitor"

func Start() {
	synctask := new(SyncTaskService)
	go synctask.StartCreateTask()
	go synctask.StartExecuteTask()
	go monitor.Start()
}
