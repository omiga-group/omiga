package configuration

import "github.com/spf13/viper"

const AppSettingsConfigKey = "app"

type AppSettings struct {
	ListeningInterface string `json:"listeningInterface"`
	Source             string `json:"source"`
}

func GetAppSettings(viper *viper.Viper) AppSettings {
	key := AppSettingsConfigKey + KeyDelimiter

	return AppSettings{
		ListeningInterface: viper.GetString(key + "listeningInterface"),
		Source:             viper.GetString(key + "source"),
	}
}
