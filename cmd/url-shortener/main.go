package main

import "flag"

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/config.yaml", "path to YAML-config file")
}

func main() {
	flag.Parse()

}
