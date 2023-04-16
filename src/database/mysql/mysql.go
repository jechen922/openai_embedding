package mysql

import (
	"fmt"
	"log"
	"openaigo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	DBOpenAI = "OpenAI"
)

var tableMap = map[string][]string{
	DBOpenAI: {"embedding"},
}

func DSNMap(cfg config.Config) map[string]string {
	dsnMap := make(map[string]string, len(tableMap))
	dsnMap[DBOpenAI] = cfg.GetMysqlENV().DSNAccount
	return dsnMap
}

func Tables(dbName string) []string {
	tables, _ := tableMap[dbName]
	return tables
}

func GormConfig(logMode logger.LogLevel) *gorm.Config {
	return &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logMode),
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	}
}

func NewConnection(cfg config.Config) IDB {
	conn = &connection{
		config: cfg,
		pool:   make(map[string]*gorm.DB),
	}
	mode := logger.Silent
	if cfg.GetSystemENV().RunMode == config.RunModeLocal {
		mode = logger.Info
	}
	gormConfig := GormConfig(mode)
	for dbName, dsn := range DSNMap(cfg) {
		gormDB, err := gorm.Open(
			mysql.New(mysql.Config{
				DSN:                       dsn,   // data source name
				DefaultStringSize:         256,   // default size for string fields
				DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
				DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
				DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
				SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
			}), gormConfig)
		if err != nil {
			log.Fatal(fmt.Sprintf("NewMysql failed at config.Conn.DB(): %v", err))
		}
		sqlDB, err := gormDB.DB()
		sqlDB.SetConnMaxIdleTime(0)
		sqlDB.SetConnMaxLifetime(0)
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(2)
		conn.pool[dbName] = gormDB
	}
	return conn
}
