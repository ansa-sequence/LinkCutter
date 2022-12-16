package ansaserver

type Config struct {
	BindAddr    string `toml:"bindAddr"`
	LogLevel    string `toml:"logLevel"`
	DatabaseURL string `toml:"databaseUrl"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "info",
	}
}
