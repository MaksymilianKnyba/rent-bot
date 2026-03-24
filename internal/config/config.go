package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
IntervalSeconds int `yaml:"interval_seconds"`
DiscordWebhook string `yaml:"discord_webhook"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
var cfg Config

err = yaml.Unmarshal(data, &cfg)
if err != nil {
	return nil, err
}

return &cfg, nil

}