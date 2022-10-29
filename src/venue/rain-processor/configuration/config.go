package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Rain     RainConfig              `yaml:"rain"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}

type RainConfig struct {
	Id                   string `yaml:"id" env:"OMIGA_RAIN_ID"`
	WebsiteUrl           string `yaml:"websiteUrl" env:"OMIGA_RAIN_WEBSITEURL"`
	BaseUrl              string `yaml:"baseUrl" env:"OMIGA_RAIN_BASEURL"`
	Username             string `yaml:"username" env:"OMIGA_RAIN_USERNAME"`
	Password             string `yaml:"password" env:"OMIGA_RAIN_PASSWORD"`
	TotpSecret           string `yaml:"totpSecret" env:"OMIGA_RAIN_TOTPSECRET"`
	Headless             bool   `yaml:"headless" env:"OMIGA_RAIN_HEADLESS"`
	Timeout              string `yaml:"timeout" env:"OMIGA_RAIN_TIMEOUT"`
	RecordedVideoDirPath string `yaml:"recordedVideoDirPath" env:"OMIGA_RAIN_RECORDEDVIDEODIRPATH"`
}
