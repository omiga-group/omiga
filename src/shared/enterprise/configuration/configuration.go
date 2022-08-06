package configuration

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func SetupConfigReader(path string) error {
	env := strings.Trim(os.Getenv("ENVIRONMENT"), " ")

	if len(env) > 0 {
		viper.SetConfigName("config." + env)
	} else {
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	trimmedPath := strings.Trim(path, " ")
	if len(trimmedPath) > 0 {
		if _, err := os.Stat(trimmedPath); !os.IsNotExist(err) {
			viper.AddConfigPath(trimmedPath)
		}
	}

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
