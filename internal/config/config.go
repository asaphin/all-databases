package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

const configFilePath = "./config.yml"

var configInstance *Config
var configInstanceSync sync.Once

func Get() *Config {
	configInstanceSync.Do(func() {
		data, err := os.ReadFile(configFilePath)
		if err != nil {
			panic(err)
		}

		var config Config

		err = yaml.Unmarshal(data, &config)
		if err != nil {
			panic(err)
		}

		configInstance = &config
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
