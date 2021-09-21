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
)

func init() {

	file, err := os.Open("config.conf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	config := config.NewConfig(buffer)
	fmt.Println(config)

	logger, err = logger2.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
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
	userRepository := repository.NewUserRepository(logger)
	userInfoRepository := repository.NewUserInfoRepository(logger)

	// UseCase
	signalingUseCase := application.NewSignalingUseCase(userRepository, userInfoRepository, logger)

	server := server.NewServer(signalingUseCase, logger)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal().Err(err)
	}
	logger.Info().Msg("Serve is running")

}
