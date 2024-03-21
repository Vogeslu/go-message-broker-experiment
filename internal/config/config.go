package config

import (
	"message_broker/internal/logger"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	} `yaml:"server"`
}

var AppConfig Config

func LoadConfig() {
	logger.Logger.Info().Msg("Loading configuration")

	file, err := os.Open("./configs/config.yml")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		panic(err)
	}
}
