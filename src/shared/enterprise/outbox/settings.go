package outbox

const ConfigKey = "outbox"

type OutboxSettings struct {
	MaxRetryCount int
	RetryInterval string
}
