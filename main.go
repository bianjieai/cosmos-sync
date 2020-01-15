package main

import (
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/tasks"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	c := make(chan os.Signal)

	defer func() {
		logger.Info("System Exit")

		models.Close()

		if err := recover(); err != nil {
			logger.Error("occur error", logger.Any("err", err))
			os.Exit(1)
		}
	}()

	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	tasks.Start()
	<-c
}
