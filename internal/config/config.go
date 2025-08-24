package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTP        HTTPConfig `yaml:"http" env-required:"true"`
}
type HTTPConfig struct {
	Port int `yaml:"port" env-required:"true"`
}

func LoadConfg(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); err != nil {
		return nil, err
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func MustLoadConfig(configPath string) *Config {
	config, err := LoadConfg(configPath)
	if err != nil {
		panic("err on parse config: " + err.Error())
	}
	return config
}
