package main

import (
	"flag"
	"fmt"
	"os/signal"
	"strings"
	"syscall"

	conf "kmp-news-consumer/config"
	log "kmp-news-consumer/logging"
	mdl "kmp-news-consumer/models"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Logf("OS: %s", runtime.GOOS)
	log.Logf("architecture: %s", runtime.GOARCH)

	// load config file
	configFile := flag.String("conf", "config/conf.yml", "main configuration file")
	flag.Parse()

	log.Logf("reads configuration from %s", *configFile)
	conf.LoadConfigFromFile(configFile)

	log.Init(conf.Param.Log.Level, conf.Param.Log.FileName)

	// initiate Service Database connection
	dbConn, err := conf.New(conf.Param.DBType, conf.Param.DBUrl)

	if err != nil {
		log.Errorf("Unable to open database %v", err)
		os.Exit(1)
	}

	// initiate Service Kafka
	kafReader := conf.KafkaConn(strings.Split(conf.Param.KafkaURL, ","), conf.Param.KafkaTopic, conf.Param.KafkaGroup)

	// Consumer
	mdl.Consumer(kafReader, dbConn, conf.Param.ElasticURL)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Logf("exit", <-errs)
}
