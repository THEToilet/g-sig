package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type config struct {
	Title 	string
	Server 	map[string]server
	LogInfo logInfo
}

type server struct {
	ServerAddress 	string 	`toml:"server_address"`
	ServerPort 		uint 	`toml:"server_port"`
}

type logInfo struct {
}

func NewConfig(buffer []byte) *config {
	var conf config
	if err := toml.Unmarshal(buffer, &conf); err != nil{
		log.Fatal(err)
	}
	return &conf
}
