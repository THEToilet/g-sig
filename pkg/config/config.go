package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Title   string
	Server  map[string]Server
	LogInfo LogInfo `toml:"log_info"`
}

type Server struct {
	ServerAddress string `toml:"server_address"`
	ServerPort    uint   `toml:"server_port"`
}

type LogInfo struct {
	Level string
}

func NewConfig(buffer []byte) *Config {
	var conf Config
	if err := toml.Unmarshal(buffer, &conf); err != nil {
		log.Fatal(err)
	}
	return &conf
}
