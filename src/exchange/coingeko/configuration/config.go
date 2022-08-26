package configuration

type Config struct {
	Coingeko CoingekoSettings `yaml:"coingeko"`
}

type CoingekoSettings struct {
	BaseUrl string `yaml:"baseUrl" env:"OMIGA_COINGEKO_BASEURL"`
}
