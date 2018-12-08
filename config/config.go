package config

import (
	"github.com/BurntSushi/toml"
)

//Config: Struct for the application coniguration
type Config struct {
	StoragePath string `toml:storage_path`
}

//ParseConfig: Function to parse config file into the config struct
func ParseConfig(file string) (*Config, error) {
	conf := &Config{}
	_, err := toml.DecodeFile(file, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
