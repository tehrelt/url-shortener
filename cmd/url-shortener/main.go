package main

import (
	"flag"
	"log"

	"github.com/tehrelt/url-shortener/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yaml", "path to YAML-config file")
}

func main() {
	flag.Parse()

	config := api.NewConfig(configPath)

	if err := api.Start(config); err != nil {
		log.Fatal(err)
	}
}
