package configuration

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func SetupConfigReader(path string) (*viper.Viper, error) {
	viperInstance := viper.NewWithOptions(viper.KeyDelimiter("_"))

	omigaEnv := strings.Trim(os.Getenv("OMIGA_ENVIRONMENT"), " ")

	if len(omigaEnv) > 0 {
		viperInstance.SetConfigName("config." + omigaEnv)
	} else {
		viperInstance.SetConfigName("config")
	}

	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath(".")

	trimmedPath := strings.Trim(path, " ")
	if len(trimmedPath) > 0 {
		if _, err := os.Stat(trimmedPath); !os.IsNotExist(err) {
			viperInstance.AddConfigPath(trimmedPath)
		}
	}

	err := viperInstance.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viperInstance.SetEnvPrefix("OMIGA")
	viperInstance.AutomaticEnv()

	return viperInstance, nil
}
