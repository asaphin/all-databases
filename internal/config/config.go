package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

const configFilePath = "./config.yml"

var configInstance *Config
var configInstanceSync sync.Once

func Get() *Config {
	configInstanceSync.Do(func() {
		log.Debug("loading config")

		data, err := os.ReadFile(configFilePath)
		if err != nil {
			log.WithError(err).Fatal("unable to read config")
		}

		var config Config

		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.WithError(err).Fatal("unable to parse config")
		}

		configInstance = &config

		log.Info("loaded config")
	})

	return configInstance
}

type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslMode"`
}
