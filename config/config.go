package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator"
)

type IConfig interface {
	GetSystemENV() SystemENV
	GetMysqlENV() MysqlENV
	GetPostgresENV() PostgresENV
	GetLoggerENV() LoggerEnv
	GetServiceName() string
}

func New() IConfig {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(fmt.Errorf("parse config fail, reason : %v", err))
	}

	if err := validator.New().Struct(&cfg); err != nil {
		log.Fatal(fmt.Errorf("validate config fail, reason : %v", err))
	}
	return cfg
}

func (cfg config) GetSystemENV() SystemENV {
	return cfg.SystemENV
}

func (cfg config) GetMysqlENV() MysqlENV {
	return cfg.MysqlENV
}

func (cfg config) GetPostgresENV() PostgresENV {
	return cfg.PostgresENV
}

func (cfg config) GetLoggerENV() LoggerEnv {
	return cfg.LoggerEnv
}

func (cfg config) GetServiceName() string {
	return serviceName
}
