package main

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/handlers"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/libs/stream"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/tasks"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
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
	config.InitEnv()
	conf, err := config.ReadConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}

	models.Init(conf)
	pool.Init(conf)
	handlers.InitMsgParser()
	if conf.Redis.Addrs != "" {
		if err = stream.Init(conf); err != nil {
			logger.Fatal(err.Error())
		}
		stream.InitMQClient(conf)
	}
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	tasks.Start(tasks.NewSyncTask(conf))
	MonitorServerStart(conf)
	<-c
}

func MonitorServerStart(conf *config.Config) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	srv := &http.Server{
		Addr:    ":" + conf.Server.PromethousPort,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
