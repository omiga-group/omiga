package configuration

type AppConfig struct {
	ListeningInterface string `yaml:"listeningInterface" env:"OMIGA_APP_LISTENINGINTERFACE"`
	Source             string `yaml:"source" env:"OMIGA_APP_SOURCE"`
}
