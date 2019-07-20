package config

import (
	"os"

	logger "kmp-news-consumer/logging"
	"kmp-news-consumer/parser"
)

// Configuration stores global configuration loaded from yml file
type Configuration struct {
	ListenPort string `yaml:"listenPort"`
	Query      string `yaml:"query"`
	DBUrl      string `yaml:"dbUrl"`
	DBType     string `yaml:"dbType"`
	Log        struct {
		FileName    string `yaml:"filename"`
		Level       string `yaml:"level"`
		KafkaOffset string `yaml:"kafkaOffset"`
	} `yaml:"log"`

	ElasticURL   string `yaml:"elasticURL"`
	ElasticIndex string `yaml:"elasticIndex"`
	KafkaURL     string `yaml:"kafkaURL"`
	KafkaTopic   string `yaml:"kafkaTopic"`
	KafkaGroup   string `yaml:"kafkaGroup"`
}

// Param use as global variable for configuration
var Param Configuration

// LoadConfigFromFile use to load global configuration
func LoadConfigFromFile(fn *string) {
	if err := parser.LoadYAML(fn, &Param); err != nil {
		logger.Errorf("LoadConfigFromFile() - Failed opening config file %s\n%s", &fn, err)
		os.Exit(1)
	}
	//logger.Logf("Loaded configs: %v", Param)
	logger.Logf("Config %s", "Loaded")
}
