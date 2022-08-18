package pulsar

type PulsarConfig struct {
	Url              string `yaml:"url", env:"OMIGA_PULSAR_URL"`
	ProducerName     string `yaml:"producerName", env:"OMIGA_PULSAR_PRODUCERNAME"`
	SubscriptionName string `yaml:"subscriptionName", env:"OMIGA_PULSAR_SUBSCRIPTIONNAME"`
}
