package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	BindIP         string `yaml:"bind_ip" env-default:"localhost"`
	Port           string `yaml:"port" env-default:"8080"`
	RabbitMQ       string `yaml:"rabbitmq"`
	Minio          string `yaml:"minio"`
	MinioAccessKey string `yaml:"minio_access_key"`
	MinioPassword  string `yaml:"minio_password"`
}

func GetConfig(environment string) (*Config, error) {

	if environment == "prod" {

		return &Config{
			BindIP:         os.Getenv("PRODUCER_HOST"),
			Port:           os.Getenv("PRODUCER_PORT"),
			RabbitMQ:       os.Getenv("RabbitMQ"),
			Minio:          os.Getenv("MINIO"),
			MinioAccessKey: os.Getenv("MINIO_ACCESS_KEY"),
			MinioPassword:  os.Getenv("MINIO_PASSWORD"),
		}, nil

	}

	instance := &Config{}
	if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
		return nil, err
	}

	return instance, nil
}
