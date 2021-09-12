package config

type config struct {
	GSigServerAddress 	string 	`toml:""`
	GSigServerPort 		uint 	`toml:""`
	WSigServerAddress 	string	`toml:""`
	WSigServerPort 		uint 	`toml:""`
}

func NewConfig(buf) *config {
	return &config{

	}
}
