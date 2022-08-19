package repositories

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/database"
	"go.uber.org/zap"
)

// 2) also why do you name the interface after the implementation?
// 3) if you plan to alway use ent go why use an interface at all!
// 4) and finally usually it is not a good practice to put interface and impementation next to each other unless you have a convincing reason
type EntgoClient interface {
	GetClient() *Client // 1) hey mori, why do you return the internal object that you are trying to abstract away, that defeats the purpose of the interface!
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

	driver := database.GetDriver()

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
