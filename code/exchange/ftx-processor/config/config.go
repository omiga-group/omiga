package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Intialize(path string, env string) error {
	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	if err = validate(env); err != nil {
		return err
	}

	return nil
}

func validate(env string) error {
	apiKey := "integration.api-key"
	if viper.GetString(apiKey) == "" {
		return fmt.Errorf("'%s' is empty", apiKey)
	}

	return nil
}
