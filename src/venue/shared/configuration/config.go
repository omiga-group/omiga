package configuration

type VenueConfig struct {
	Id string `yaml:"id" env:"OMIGA_VENUE_ID"`
}
