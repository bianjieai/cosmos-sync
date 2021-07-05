package tasks

func Start(synctask SyncTask) {
	go synctask.StartCreateTask()
	go synctask.StartExecuteTask()
}
