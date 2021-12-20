package main

import (
	"flag"
	"fmt"
	"g-sig/pkg/config"
	"g-sig/pkg/domain/application"
	"g-sig/pkg/gateway/repository"
	logger2 "g-sig/pkg/logger"
	"g-sig/pkg/server"
	"github.com/rs/zerolog"
	"io/ioutil"
	"log"
	"os"
)

var (
	version = "0.1.0"
	logger  *zerolog.Logger
	con     *config.Config
)

func init() {

	file, err := os.Open("g-sig.conf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	con = config.NewConfig(buffer)
	fmt.Println(con)

	logger, err = logger2.NewLogger(con)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info().Str("Title", con.Title).Msg("Config")
	logger.Info().Str("LogLevel", con.LogInfo.Level).Msg("Config")
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Printf("g-sig version is %s", version)
		return
	}

	// Repository
	userRepository := repository.NewUserRepository()
	userInfoRepository := repository.NewUserInfoRepository()

	// UseCase
	signalingUseCase := application.NewSignalingUseCase(userRepository, userInfoRepository, logger)

	server := server.NewServer(signalingUseCase, logger)

	logger.Info().Str("Addr", ":8080").Msg("Serve is running")

	if _, err := os.Stat(con.Cert); err == nil {
		if err := server.ListenAndServeTLS(con.Cert, con.Key); err != nil {
			logger.Fatal().Err(err).Msg("wss error")
		}
	} else {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal().Err(err).Msg("ws error")
		}
	}
}
