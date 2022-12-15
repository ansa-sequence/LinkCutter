package store

type Config struct {
	DatabaseURL string `toml:"databaseUrl"`
}

func NewConfig() *Config {
	return &Config{}
}
