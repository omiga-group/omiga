package postgres

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/omiga-group/omiga/src/shared/enterprise/database"
	"go.uber.org/zap"
)

type postgres struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

type brokenConnectiongString struct {
	databaseName string
	host         string
	port         string
	username     string
	password     string
}

func NewPostgres(
	logger *zap.SugaredLogger,
	postgresConfig PostgresConfig) (database.Database, error) {
	brokenConnectiongString, err := parseConnectionString(postgresConfig.ConnectionString)
	if err != nil {
		return nil, err
	}

	logger.Infof(
		"Setting up Postgres connection string to connect to Postgres: host: %s, port: %s, databaseName: %s",
		brokenConnectiongString.host,
		brokenConnectiongString.port,
		brokenConnectiongString.databaseName)

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

func (p *postgres) GetDriver() *entsql.Driver {
	return entsql.OpenDB(dialect.Postgres, p.db)
}

func (p *postgres) GetDB() *sql.DB {
	return p.db
}

func (p *postgres) Close() {
	if p.db != nil {
		if err := p.db.Close(); err != nil {
			p.logger.Errorf("Failed to close database. Error: %v", err)
		}

		p.db = nil
	}
}

func parseConnectionString(connectionString string) (brokenConnectiongString, error) {
	result := brokenConnectiongString{}

	u, err := url.Parse(connectionString)
	if err != nil {
		return result, err
	}

	if u.Scheme != "postgres" && u.Scheme != "postgresql" {
		return result, fmt.Errorf("invalid connection protocol: %s", u.Scheme)
	}

	if u.User != nil {
		v := u.User.Username()
		if v != "" {
			result.username = v
		}

		v, _ = u.User.Password()
		if v != "" {
			result.password = v
		}
	}

	if host, port, err := net.SplitHostPort(u.Host); err != nil {
		if u.Host != "" {
			result.host = u.Host
		}
	} else {
		if host != "" {
			result.host = host
		}

		if port != "" {
			result.port = port
		}
	}

	if u.Path != "" {
		if u.Path[1:] != "" {
			result.databaseName = u.Path[1:]
		}
	}

	return result, nil
}
