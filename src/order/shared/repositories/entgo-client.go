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
	logger *zap.SugaredLogger
	client *Client
}

func NewEntgoClient(
	logger *zap.SugaredLogger,
	database database.Database) (EntgoClient, error) {

	driver, err := database.GetDriver()
	if err != nil {
		return nil, err
	}

	return &entgoClient{
		logger: logger,
		client: NewClient(Driver(driver)),
	}, nil
}

func (ec *entgoClient) GetClient() (*Client, error) {
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
