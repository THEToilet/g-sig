package application

import (
	"g-sig/pkg/domain/model"
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

func (s *SignalingUseCase) Register(userID string, location model.GeoLocation) error {
	user := model.NewUserInfo(userID, location.Latitude, location.Longitude)
	s.logger.Info().Msg(userID)
	return s.userInfoRepository.Save(*user)
}

func (s *SignalingUseCase) Update(userInfo model.UserInfo) error {
	return s.userInfoRepository.Update(userInfo)
}

func (s *SignalingUseCase) Delete(userInfo model.UserInfo) error {
	return s.userInfoRepository.Delete(userInfo.UserID)
}

func (s *SignalingUseCase) StaticSearch(userInfo model.UserInfo, searchDistance float64) []*model.UserInfo {
	userInfoList := s.userInfoRepository.FindAll()
	var searchedUserList []*model.UserInfo
	for _, v := range userInfoList {
		if (((userInfo.Latitude - v.Latitude) * (userInfo.Latitude - v.Latitude)) + ((userInfo.Longitude - v.Longitude) * (userInfo.Longitude - v.Longitude))) <= searchDistance*searchDistance {
			if v.UserID != userInfo.UserID {
				searchedUserList = append(searchedUserList, v)
			}
		}
	}
	return searchedUserList
}

func (s *SignalingUseCase) DynamicSearch(userInfo model.UserInfo, searchDistance float64) []*model.UserInfo {
	userInfoList := s.userInfoRepository.FindAll()
	// DoSomething
	return userInfoList
}

func (s *SignalingUseCase) Send() {
}
