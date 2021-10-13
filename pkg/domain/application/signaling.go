package application

import (
	"g-sig/pkg/domain/model"
	"g-sig/pkg/domain/repository"
	"github.com/rs/zerolog"
	"math"
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
	return s.userInfoRepository.Save(*user)
}

func (s *SignalingUseCase) Update(userInfo model.UserInfo) error {
	return s.userInfoRepository.Update(userInfo)
}

func (s *SignalingUseCase) Delete(userID string) error {
	return s.userInfoRepository.Delete(userID)
}

func (s *SignalingUseCase) StaticSearch(userInfo model.UserInfo, searchDistance float64) []*model.UserInfo {
	userInfoList := s.userInfoRepository.FindAll()
	var searchedUserList []*model.UserInfo
	for _, v := range userInfoList {
		//fmt.Println ("2点間", userInfo.Latitude, userInfo.Longitude, v.Latitude, v.Longitude)
		//fmt.Println("2点間の距離", 1000*6371*math.Acos(math.Cos(userInfo.Latitude*math.Pi/180)*math.Cos(v.Latitude*math.Pi/180)*math.Cos(v.Longitude*math.Pi/180-userInfo.Longitude*math.Pi/180)+math.Sin(userInfo.Latitude*math.Pi/180)*math.Sin(v.Latitude*math.Pi/180)))
		if 1000*6371*math.Acos(math.Cos(userInfo.Latitude*math.Pi/180)*math.Cos(v.Latitude*math.Pi/180)*math.Cos(v.Longitude*math.Pi/180-userInfo.Longitude*math.Pi/180)+math.Sin(userInfo.Latitude*math.Pi/180)*math.Sin(v.Latitude*math.Pi/180)) <= searchDistance {
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
