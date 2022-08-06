package repositories

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/database"
	"go.uber.org/zap"
)

type EntgoClient interface {
	GetClient() (*Client, error)
	Close()
}

type entgoClient struct {
	logger   *zap.SugaredLogger
	database database.Database
	client   *Client
}

func NewEntgoClient(
	logger *zap.SugaredLogger,
	database database.Database) (EntgoClient, error) {

	return &entgoClient{
		logger:   logger,
		database: database,
	}, nil
}

func (ec *entgoClient) GetClient() (*Client, error) {
	if ec.client != nil {
		return ec.client, nil
	}

	driver, err := ec.database.GetDriver()
	if err != nil {
		return nil, err
	}

	ec.client = NewClient(Driver(driver))

	return ec.client, nil
}

func (ec *entgoClient) Close() {
	if ec.client != nil {
		if err := ec.client.Close(); err != nil {
			ec.logger.Errorf("Failed to close entgo client. Error: %v", err)
		}

		ec.client = nil
	}
}
