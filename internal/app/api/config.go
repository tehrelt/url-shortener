package api

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port             string `yaml:"port"`
	Env              string `yaml:"env"`
	ConnectionString string `yaml:"connection-string"`
}

func NewConfig(configPath string) *Config {

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatal(err)
	}

	return &config
}
