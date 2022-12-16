package main

import (
	"LinkCutter/internal/ansaserver"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/ansaserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := ansaserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := ansaserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
