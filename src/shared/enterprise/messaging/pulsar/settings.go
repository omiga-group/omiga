package pulsar

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/spf13/viper"
)

const ConfigKey = "pulsar"

type PulsarSettings struct {
	Url              string `json:"url"`
	ProducerName     string `json:"producerName"`
	SubscriptionName string `json:"subscriptionName"`
}

func GetPulsarSettings(viper *viper.Viper) PulsarSettings {
	key := ConfigKey + configuration.KeyDelimiter

	return PulsarSettings{
		Url:              viper.GetString(key + "url"),
		ProducerName:     viper.GetString(key + "producerName"),
		SubscriptionName: viper.GetString(key + "subscriptionName"),
	}
}
