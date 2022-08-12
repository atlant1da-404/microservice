package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	RabbitMQ       string `yaml:"rabbitmq"`
	Minio          string `yaml:"minio"`
	MinioAccessKey string `yaml:"minio_access_key"`
	MinioPassword  string `yaml:"minio_password"`
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
