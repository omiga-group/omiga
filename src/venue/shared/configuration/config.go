package configuration

type ExchangeConfig struct {
	Id string `yaml:"id" env:"OMIGA_EXCHANGE_ID"`
}
