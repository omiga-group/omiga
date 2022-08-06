package postgres

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/omiga-group/omiga/src/shared/enterprise/database"
)

type PostgresSettings struct {
	ConnectionString string
	MaxOpenConns     int
}

type postgres struct {
	postgresSettings PostgresSettings
}

func NewPostgres(
	postgresSettings PostgresSettings) (database.Database, error) {

	return &postgres{
		postgresSettings: postgresSettings,
	}, nil
}

func (p *postgres) GetDriver() (*entsql.Driver, error) {
	db, err := sql.Open("pgx", p.postgresSettings.ConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(p.postgresSettings.MaxOpenConns)

	return entsql.OpenDB(dialect.Postgres, db), nil
}
