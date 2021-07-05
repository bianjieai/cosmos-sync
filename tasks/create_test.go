package tasks

import (
	"testing"
	"time"
)

func TestSyncTaskService_createTask(t *testing.T) {
	s := syncTaskService{}
	chanLimit := make(chan bool, 1)

	for {
		chanLimit <- true
		go s.createTask(100, chanLimit)
		time.Sleep(time.Duration(1) * time.Minute)
	}
}
