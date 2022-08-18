package postgres

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/omiga-group/omiga/src/shared/enterprise/database"
	"go.uber.org/zap"
)

type postgres struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewPostgres(
	logger *zap.SugaredLogger,
	postgresConfig PostgresConfig) (database.Database, error) {
	db, err := sql.Open("pgx", postgresConfig.ConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(postgresConfig.MaxOpenConns)

	return &postgres{
		db:     db,
		logger: logger,
	}, nil
}

func (p *postgres) GetDriver() (*entsql.Driver, error) {
	return entsql.OpenDB(dialect.Postgres, p.db), nil
}

func (p *postgres) GetDB() (*sql.DB, error) {
	return p.db, nil
}

func (p *postgres) Close() {
	if p.db != nil {
		if err := p.db.Close(); err != nil {
			p.logger.Errorf("Failed to close database. Error: %v", err)
		}

		p.db = nil
	}
}
