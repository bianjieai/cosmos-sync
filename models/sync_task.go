package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNameSyncTask = "sync_task"

	// value of status
	SyncTaskStatusUnHandled = "unhandled"
	SyncTaskStatusUnderway  = "underway"
	SyncTaskStatusCompleted = "completed"

	// only for follow task
	// when current_height of follow task add blockNumPerWorkerHandle
	// less than blockchain current_height, this follow task's status should be set invalid
	FollowTaskStatusInvalid = "invalid"

	// taskType
	SyncTaskTypeCatchUp = "catch_up"
	SyncTaskTypeFollow  = "follow"
)

type (
	SyncTask struct {
		ID             bson.ObjectId `bson:"_id"`
		StartHeight    int64         `bson:"start_height"`     // task start height
		EndHeight      int64         `bson:"end_height"`       // task end height
		CurrentHeight  int64         `bson:"current_height"`   // task current height
		Status         string        `bson:"status"`           // task status
		WorkerId       string        `bson:"worker_id"`        // worker id
		WorkerLogs     []WorkerLog   `bson:"worker_logs"`      // worker logs
		LastUpdateTime int64         `bson:"last_update_time"` // unix timestamp
	}

	WorkerLog struct {
		WorkerId  string    `bson:"worker_id"`  // worker id
		BeginTime time.Time `bson:"begin_time"` // time which worker begin handle this task
	}
)

func (d SyncTask) Name() string {
	if GetSrvConf().ChainId == "" {
		return CollectionNameSyncTask
	}
	return fmt.Sprintf("sync_%v_task", GetSrvConf().ChainId)
}

func (d SyncTask) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-start_height", "-end_height"},
		Unique:     true,
		Background: true,
	})
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-status"},
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}

func (d SyncTask) PkKvPair() map[string]interface{} {
	return bson.M{"start_height": d.CurrentHeight, "end_height": d.EndHeight}
}

// get max block height in sync task
func (d SyncTask) GetMaxBlockHeight() (int64, error) {
	type maxHeightRes struct {
		MaxHeight int64 `bson:"max"`
	}
	var res []maxHeightRes

	q := []bson.M{
		{
			"$group": bson.M{
				"_id": nil,
				"max": bson.M{"$max": "$end_height"},
			},
		},
	}

	getMaxBlockHeightFn := func(c *mgo.Collection) error {
		return c.Pipe(q).All(&res)
	}
	err := ExecCollection(d.Name(), getMaxBlockHeightFn)

	if err != nil {
		return 0, err
	}
	if len(res) > 0 {
		return res[0].MaxHeight, nil
	}

	return 0, nil
}

// query record by status
func (d SyncTask) QueryAll(status []string, taskType string) ([]SyncTask, error) {
	var syncTasks []SyncTask
	q := bson.M{}

	if len(status) > 0 {
		q["status"] = bson.M{
			"$in": status,
		}
	}

	switch taskType {
	case SyncTaskTypeCatchUp:
		q["end_height"] = bson.M{
			"$ne": 0,
		}
		break
	case SyncTaskTypeFollow:
		q["end_height"] = bson.M{
			"$eq": 0,
		}
		break
	}

	fn := func(c *mgo.Collection) error {
		return c.Find(q).All(&syncTasks)
	}

	err := ExecCollection(d.Name(), fn)

	if err != nil {
		return syncTasks, err
	}

	return syncTasks, nil
}

func (d SyncTask) GetExecutableTask(maxWorkerSleepTime int64) ([]SyncTask, error) {
	var tasks []SyncTask

	t := time.Now().Add(time.Duration(-maxWorkerSleepTime) * time.Second).Unix()
	q := bson.M{
		"status": bson.M{
			"$in": []string{SyncTaskStatusUnHandled, SyncTaskStatusUnderway},
		},
	}

	fn := func(c *mgo.Collection) error {
		return c.Find(q).Sort("-status").Limit(1000).All(&tasks)
	}

	err := ExecCollection(d.Name(), fn)

	if err != nil {
		return tasks, err
	}

	ret := make([]SyncTask, 0, len(tasks))
	//filter the task which last_update_time >= now
	for _, task := range tasks {
		if task.LastUpdateTime >= t && task.Status == SyncTaskStatusUnderway {
			continue
		}
		ret = append(ret, task)
	}

	return ret, nil
}

func (d SyncTask) GetTaskById(id bson.ObjectId) (SyncTask, error) {
	var task SyncTask

	fn := func(c *mgo.Collection) error {
		return c.FindId(id).One(&task)
	}

	err := ExecCollection(d.Name(), fn)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (d SyncTask) GetTaskByIdAndWorker(id bson.ObjectId, worker string) (SyncTask, error) {
	var task SyncTask

	fn := func(c *mgo.Collection) error {
		q := bson.M{
			"_id":       id,
			"worker_id": worker,
		}

		return c.Find(q).One(&task)
	}

	err := ExecCollection(d.Name(), fn)
	if err != nil {
		return task, err
	}
	return task, nil
}

// take over a task
// update status, worker_id, worker_logs and last_update_time
func (d SyncTask) TakeOverTask(task SyncTask, workerId string) error {
	// multiple goroutine attempt to update same record,
	// use this selector to ensure only one goroutine can update success at same time
	fn := func(c *mgo.Collection) error {
		selector := bson.M{
			"_id":              task.ID,
			"last_update_time": task.LastUpdateTime,
		}

		task.Status = SyncTaskStatusUnderway
		task.WorkerId = workerId
		task.LastUpdateTime = time.Now().Unix()
		task.WorkerLogs = append(task.WorkerLogs, WorkerLog{
			WorkerId:  workerId,
			BeginTime: time.Now(),
		})

		return c.Update(selector, task)
	}

	return ExecCollection(d.Name(), fn)
}

// update task last update time
func (d SyncTask) UpdateLastUpdateTime(task SyncTask) error {
	fn := func(c *mgo.Collection) error {
		selector := bson.M{
			"_id":       task.ID,
			"worker_id": task.WorkerId,
		}

		task.LastUpdateTime = time.Now().Unix()

		return c.Update(selector, task)
	}

	return ExecCollection(d.Name(), fn)
}

// query valid follow way
func (d SyncTask) QueryValidFollowTasks() (bool, error) {
	var syncTasks []SyncTask
	q := bson.M{}

	q["status"] = SyncTaskStatusUnderway

	q["end_height"] = bson.M{
		"$eq": 0,
	}

	fn := func(c *mgo.Collection) error {
		return c.Find(q).All(&syncTasks)
	}

	err := ExecCollection(d.Name(), fn)

	if err != nil {
		return false, err
	}

	if len(syncTasks) == 1 {
		return true, nil
	}

	return false, nil
}
