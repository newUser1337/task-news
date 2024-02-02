package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Core Core `envconfig:"core"`
}

type Core struct {
	Port            int           `envconfig:"port"`
	RefreshInterval time.Duration `envconfig:"refresh_interval"`
	Mongo           *Mongo        `envconfig:"mongo"`
}

type Mongo struct {
	Address   string `envconfig:"address"`
	TableName string `envconfig:"tablename"`
}

var configuration *Config

func init() {
	configuration = &Config{}
	if err := envconfig.Process("TASK_NEWS", configuration); err != nil {
		log.Fatalf("failed to read config %s", err)
	}
}

func GetConfig() *Config {
	return configuration
}
