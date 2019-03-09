package app

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug   bool   `toml:"debug" envconfig:"debug"`
	Address string `toml:"address" envconfig:"address"`
}

type App struct {
	Config *Config
}

func NewApp(path string) (*App, error) {
	cfg, err := newConfig(path)
	if err != nil {
		return nil, err
	}

	app := &App{
		Config: cfg,
	}
	return app, nil
}

func newConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	if _, err := toml.DecodeReader(f, &config); err != nil {
		return nil, err
	}

	if err := envconfig.Process("sample", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
