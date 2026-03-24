package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
InternalSeconds int `yaml:"internal_seconds"`
DiscordWebhook string `yaml:"discord_webhook"`
}

func Load(path string) (*Config, error) {
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