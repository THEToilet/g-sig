package application

import (
	"g-sig/pkg/domain/repository"
	"github.com/rs/zerolog"
)

type SignalingUseCase struct {
	userRepository     repository.UserRepository
	userInfoRepository repository.UserInfoRepository
	logger             *zerolog.Logger
}

func NewSignalingUseCase(userRepository repository.UserRepository, userInfoRepository repository.UserInfoRepository, logger *zerolog.Logger) *SignalingUseCase {
	return &SignalingUseCase{
		userRepository:     userRepository,
		userInfoRepository: userInfoRepository,
		logger:             logger,
	}
}

func (s *SignalingUseCase) Register() {
	s.userInfoRepository.Save()
}

func (s *SignalingUseCase) Update() {
	s.userInfoRepository.Update()
}

func (s *SignalingUseCase) Delete() {
	s.userInfoRepository.Delete()
}

func (s *SignalingUseCase) StaticSearch() {
	userInfoList, err :=  s.userInfoRepository.FindAll()
}

func (s *SignalingUseCase) DynamicSearch() {
	userInfoList, err :=  s.userInfoRepository.FindAll()
}

func (s *SignalingUseCase) Send() {
}
