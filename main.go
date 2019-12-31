package main

import (
	"gitlab.bianjie.ai/irita/ex-sync/libs/logger"
	"gitlab.bianjie.ai/irita/ex-sync/models"
	"gitlab.bianjie.ai/irita/ex-sync/tasks"
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
