package repositories

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/database"
	"go.uber.org/zap"
)

type EntgoClient interface {
	GetClient() *Client
	Close()
}

type entgoClient struct {
	logger   *zap.SugaredLogger
	client   *Client
	database database.Database
}

func NewEntgoClient(
	logger *zap.SugaredLogger,
	database database.Database) (EntgoClient, error) {

	driver, err := database.GetDriver()
	if err != nil {
		return nil, err
	}

	return &entgoClient{
		logger:   logger,
		client:   NewClient(Driver(driver)),
		database: database,
	}, nil
}

func (ec *entgoClient) GetClient() *Client {
	return ec.client
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
