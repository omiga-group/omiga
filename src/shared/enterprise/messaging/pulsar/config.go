package pulsar

type PulsarConfig struct {
	Url                        string                            `yaml:"url" env:"OMIGA_PULSAR_URL"`
	ProducerName               string                            `yaml:"producerName" env:"OMIGA_PULSAR_PRODUCERNAME"`
	SubscriptionName           string                            `yaml:"subscriptionName" env:"OMIGA_PULSAR_SUBSCRIPTIONNAME"`
	OperationTimeout           string                            `yaml:"operationTimeout" env:"OMIGA_PULSAR_OPERATIONTIMEOUT"`
	ConnectionTimeout          string                            `yaml:"connectionTimeout" env:"OMIGA_PULSAR_CONNECTIONTIMEOUT"`
	EnableAuthenticationOAuth2 bool                              `yaml:"enableAuthenticationOAuth2" env:"OMIGA_PULSAR_ENABLEAUTHENTICATIONOAUTH2"`
	AuthenticationOAuth2       PulsarAuthenticationOAuth2Config  `yaml:"authenticationOAuth2,omitempty"`
}

type PulsarAuthenticationOAuth2Config struct {
	Type       string `yaml:"type" env:"OMIGA_PULSAR_AUTHENTICATION_TYPE"`
	IssuerUrl  string `yaml:"issuerUrl" env:"OMIGA_PULSAR_AUTHENTICATION_URL"`
	Audience   string `yaml:"audience" env:"OMIGA_PULSAR_AUTHENTICATION_AUDIENCE"`
	PrivateKey string `yaml:"privateKey" env:"OMIGA_PULSAR_AUTHENTICATION_PRIVATEKEY"`
	ClientId   string `yaml:"clientId" env:"OMIGA_PULSAR_AUTHENTICATION_CLIENTID"`
}
