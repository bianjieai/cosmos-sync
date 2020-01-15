package tasks

func Start() {
	synctask := new(SyncTaskService)
	go synctask.StartCreateTask()
	go synctask.StartExecuteTask()
}
