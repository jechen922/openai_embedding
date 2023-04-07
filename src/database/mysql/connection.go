package mysql

import "gorm.io/gorm"

var conn *connection

type Connections struct {
	OpenAI *gorm.DB
}

type IDB interface {
	Conn(name string) *gorm.DB // 取得 DB
	Pool(dbNames ...string) *Connections
	Commit(transactions ...*gorm.DB) error
}

type connection struct {
	pool map[string]*gorm.DB
}

func (c *connection) Conn(dbName string) *gorm.DB {
	gormDB, exist := c.pool[dbName]
	if !exist || c == nil {
		return nil
	}
	return gormDB
}

func (c *connection) Pool(dbNames ...string) *Connections {
	connections := &Connections{}
	for _, dbName := range dbNames {
		switch dbName {
		case DBOpenAI:
			connections.OpenAI = c.Conn(dbName)
		}
	}
	return connections
}

func (c *connection) Commit(transactions ...*gorm.DB) error {
	for _, tx := range transactions {
		if err := tx.Commit().Error; err != nil {
			return err
		}
	}
	return nil
}
