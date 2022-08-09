package database

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
)

type Database interface {
	GetDB() (*sql.DB, error)
	GetDriver() (*entsql.Driver, error)
	Close()
}
