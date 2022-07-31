package configurations

type GeneralSettings struct {
	Port int
}

type PulsarSettings struct {
	Url string
}

type PostgresSettings struct {
	ConnectionString string
}
