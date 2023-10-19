package urlshortener

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

func NewConfig(configPath string) *Config {

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config: %w", err)
	}

	return &config
}
