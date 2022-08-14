package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	RabbitMQ       string `yaml:"rabbitmq"`
	Minio          string `yaml:"minio"`
	MinioAccessKey string `yaml:"minio_access_key"`
	MinioPassword  string `yaml:"minio_password"`
}

func GetConfig() (*Config, error) {

	instance := &Config{}
	if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
		return nil, err
	}

	return instance, nil
}
