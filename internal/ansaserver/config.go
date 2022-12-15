package ansaserver

import "LinkCutter/internal/store"

type Config struct {
	BindAddr string `toml:"bindAddr"`
	LogLevel string `toml:"logLevel"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "info",
		Store:    store.NewConfig(),
	}
}
