package database

import "entgo.io/ent/dialect/sql"

type Database interface {
	GetDriver() (*sql.Driver, error)
}
