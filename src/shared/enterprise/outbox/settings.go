package outbox

import (
	"time"

	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/spf13/viper"
)

const ConfigKey = "outbox"

type OutboxSettings struct {
	MaxRetryCount int           `json:"maxRetryCount"`
	RetryDelay    time.Duration `json:"retryDelay"`
}

func GetOutboxSettings(viper *viper.Viper) OutboxSettings {
	key := ConfigKey + configuration.KeyDelimiter

	return OutboxSettings{
		MaxRetryCount: viper.GetInt(key + "maxRetryCount"),
		RetryDelay:    viper.GetDuration(key + "retryDelay"),
	}
}
