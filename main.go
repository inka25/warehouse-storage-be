package main

import (
	"InkaTry/warehouse-storage-be/cmd/webservice"
	"InkaTry/warehouse-storage-be/internal/pkg/config"
	"InkaTry/warehouse-storage-be/internal/pkg/logger"
	"gopkg.in/ini.v1"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// if in the future some random number is needed, it needs to be seeded
	// especially for security issue
	rand.Seed(time.Now().UnixNano())

	// load config
	cfg := loadAppConfig()

	// custom log message
	location, err := time.LoadLocation(cfg.Location)
	if err != nil {
		panic(err)
	}
	log.SetOutput(&logger.LogWriter{
		AppName: cfg.AppName,
		Loc:     location,
		Env:     cfg.Env,
	})
	log.SetFlags(0)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGSTOP)

	go webservice.Start(&cfg)

	<-ch

}

func loadAppConfig() config.Config {
	appConfig := config.Config{}
	source := "./config/config.ini"
	if err := ini.MapTo(&appConfig, source); err != nil {
		panic(err)
	}
	return appConfig

}
