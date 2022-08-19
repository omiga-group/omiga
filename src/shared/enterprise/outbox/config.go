package outbox

import (
	"time"
)

type OutboxConfig struct {
	MaxRetryCount int           `yaml:"maxRetryCount" env:"OMIGA_OUTBOX_MAXRETRYCOUNT"`
	RetryDelay    time.Duration `yaml:"retryDelay" env:"OMIGA_OUTBOX_RETRYDELAY"`
}
