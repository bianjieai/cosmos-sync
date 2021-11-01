package main

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/handlers"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/tasks"
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
	conf, err := config.ReadConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}
	models.Init(conf)
	pool.Init(conf)
	handlers.InitRouter(conf)

	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	tasks.Start(tasks.NewSyncTask(conf))
	<-c
}
