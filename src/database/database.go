package database

import (
	"database/sql"
)

type (
	IDatabase interface {
		Postgres() *sql.DB
	}
	database struct {
		postgresDB *sql.DB
	}
)

func New(DB *sql.DB) IDatabase {
	return &database{
		postgresDB: DB,
	}
}

func (d database) Postgres() *sql.DB {
	return d.postgresDB
}
