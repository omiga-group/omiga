package repositories

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/database"
	"go.uber.org/zap"
)

type EntgoClient interface {
	Close()
	GetClient() *Client
	CreateTransaction(ctx context.Context) (*Tx, error)
	RollbackTransaction(tx *Tx) error
	CommitTransaction(tx *Tx) error
}

type entgoClient struct {
	logger   *zap.SugaredLogger
	client   *Client
	database database.Database
}

func NewEntgoClient(
	logger *zap.SugaredLogger,
	database database.Database) (EntgoClient, error) {

	driver := database.GetDriver()

	return &entgoClient{
		logger:   logger,
		client:   NewClient(Driver(driver)),
		database: database,
	}, nil
}

func (ec *entgoClient) Close() {
	if ec.client != nil {
		if err := ec.client.Close(); err != nil {
			ec.logger.Errorf("Failed to close entgo client. Error: %v", err)
		}

		ec.client = nil
	}

	ec.database.Close()
}

func (ec *entgoClient) GetClient() *Client {
	return ec.client
}

func (ec *entgoClient) CreateTransaction(ctx context.Context) (*Tx, error) {
	return ec.GetClient().Tx(ctx)
}

func (ec *entgoClient) RollbackTransaction(tx *Tx) error {
	return tx.Rollback()
}

func (ec *entgoClient) CommitTransaction(tx *Tx) error {
	return tx.Commit()
}
