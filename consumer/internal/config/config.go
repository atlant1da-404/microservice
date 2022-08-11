package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	RabbitMQ      string `yaml:"rabbitmq"`
	Redis         string `yaml:"redis"`
	RedisPassword string `yaml:"redis_password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatalln(help)
		}
	})
	return instance
}
