package main

import (
	"flag"
	"log"

	urlshortener "github.com/tehrelt/url-shortener/internal/url-shortener"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yaml", "path to YAML-config file")
}

func main() {
	flag.Parse()

	config := urlshortener.NewConfig(configPath)

	log.Println("init config completed")
	log.Println(config)

	server := urlshortener.New(config)
	server.Logger.Info("info log")
	server.Logger.Debug("debug log")
	server.Logger.Error("error log")

	server.Start()
}
