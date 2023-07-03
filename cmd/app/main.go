package main

import (
	"flag"
	"streamingservice/pkg/config"
	"streamingservice/pkg/server"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configpath", "config/config.toml", "path to config")
}

func main() {
	flag.Parse()

	config := config.New()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		panic(err)
	}

	s, err := server.New(*config)
	if err != nil {
		panic(err)
	}
	if err := s.Start(); err != nil {
		panic(err)
	}
}
