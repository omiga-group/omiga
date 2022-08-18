package configuration

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	App    configuration.AppConfig `yaml:"app"`
	Pulsar pulsar.PulsarConfig     `yaml:"pulsar"`
}
