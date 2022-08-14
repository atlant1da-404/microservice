package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen struct {
		BindIP string `yaml:"bind_ip" env-default:"localhost"`
		Port   string `yaml:"port" env-default:"8080"`
	}
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
