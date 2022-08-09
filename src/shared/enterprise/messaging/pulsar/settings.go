package pulsar

const ConfigKey = "pulsar"

type PulsarSettings struct {
	Url              string
	ProducerName     string
	SubscriptionName string
}
