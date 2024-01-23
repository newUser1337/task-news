package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	TaskNews `mapstructure:"tasknews"`
}

type TaskNews struct {
	Port            int           `mapstructure:"port"`
	RefreshInterval time.Duration `mapstructure:"refreshinterval"`
	Mongo           Mongo         `mapstructure:"mongo"`
}

type Mongo struct {
	Address string `mapstructure:"address"`
}

func init() {
	viper.SetConfigName("dev-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/task-news/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config %s", err)
	}
}

func GetConfig() (*Config, error) {
	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("failed to unmarshal config %s", err)
		return nil, err
	}
	return config, nil
}
