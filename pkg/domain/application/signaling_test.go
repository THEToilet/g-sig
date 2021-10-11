package application

import (
	"fmt"
	"g-sig/pkg/config"
	"g-sig/pkg/gateway/repository"
	logger2 "g-sig/pkg/logger"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestSignalingUseCase_Register(t *testing.T) {
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

	logger, err := logger2.NewLogger(config)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info().Str("Title", config.Title).Msg("Config")
	logger.Info().Str("LogLevel", config.LogInfo.Level).Msg("Config")
	// Repository
	userRepository := repository.NewUserRepository(logger)
	userInfoRepository := repository.NewUserInfoRepository(logger)

	// UseCase
	signalingUseCase := NewSignalingUseCase(userRepository, userInfoRepository, logger)
}

func TestSignalingUseCase_Update(t *testing.T) {

}

func TestSignalingUseCase_StaticSearch(t *testing.T) {

}

func TestSignalingUseCase_DynamicSearch(t *testing.T) {

}

func TestSignalingUseCase_Delete(t *testing.T) {

}

func TestSignalingUseCase_Send(t *testing.T) {

}
