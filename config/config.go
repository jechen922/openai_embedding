package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator"
)

func New() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(fmt.Errorf("parse config fail, reason : %v", err))
	}

	if err := validator.New().Struct(&cfg); err != nil {
		log.Fatal(fmt.Errorf("validate config fail, reason : %v", err))
	}
	return cfg
}

func (cfg Config) GetSystemENV() SystemENV {
	return cfg.SystemENV
}

func (cfg Config) GetPostgresENV() MysqlENV {
	return cfg.MysqlENV
}

func (cfg Config) GetServiceName() string {
	return serviceName
}
