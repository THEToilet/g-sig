package application

import (
	"g-sig/pkg/domain/model"
	"g-sig/pkg/domain/repository"
	"github.com/rs/zerolog"
	"github.com/google/uuid"
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

func (s *SignalingUseCase) Register(userInfo model.UserInfo) error {
	userID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	user := model.NewUserInfo(userID.String(), userInfo.PrivateIP, userInfo.PrivatePort, userInfo.PublicIP, userInfo.PublicPort, userInfo.Latitude, userInfo.Longitude)
	return s.userInfoRepository.Save(*user)
}

func (s *SignalingUseCase) Update(userInfo model.UserInfo) error {
	return s.userInfoRepository.Update(userInfo)
}

func (s *SignalingUseCase) Delete(userInfo model.UserInfo) error {
	return s.userInfoRepository.Delete(userInfo.UserID)
}

func (s *SignalingUseCase) StaticSearch(userInfo model.UserInfo, searchDistance float64) ([]*model.UserInfo, error){
	userInfoList, err :=  s.userInfoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var searchedUserList []*model.UserInfo
	for _, v  := range userInfoList {
		if ((userInfo.Latitude - v.Latitude) * (userInfo.Latitude - v.Latitude) + ((userInfo.Longitude- v.Longitude) * (userInfo.Longitude - v.Longitude))) <= searchDistance*searchDistance {
			if v.UserID != userInfo.UserID {
				searchedUserList = append(searchedUserList, v)
			}
		}
	}
	return searchedUserList, err
}

func (s *SignalingUseCase) DynamicSearch(userInfo model.UserInfo, searchDistance float64) ([]*model.UserInfo, error){
	userInfoList, err :=  s.userInfoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	// DoSomething
	return userInfoList, err
}

func (s *SignalingUseCase) Send() {
}
