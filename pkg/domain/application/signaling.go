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

// StaticSearch TODO: リファクタリング
func (s *SignalingUseCase) StaticSearch(userID string, geoLocation model.GeoLocation, searchDistance float64) []*model.UserInfo {
	userInfoList := s.userInfoRepository.FindAll()
	// NOTE: var searchedUserList []*model.UserInfo だとjson.Marshallでjsonがnullになるので変更した
	searchedUserList := make([]*model.UserInfo, 0)
	for _, v := range userInfoList {
		//s.logger.Debug().Interface("my-x", geoLocation.Latitude).Interface("my-y", geoLocation.Longitude).Interface("opponent-x", v.GeoLocation.Latitude).Interface("opponent-y", v.GeoLocation.Longitude).Msg("TwoPoints")
		//s.logger.Debug().Interface("2点間の距離", s.TwoPointsDistance(geoLocation, v.GeoLocation)).Msg("Distance")
		if s.TwoPointsDistance(geoLocation, v.GeoLocation) <= searchDistance {
			if v.UserID != userID {
				searchedUserList = append(searchedUserList, v)
			}
		}
	}
	return searchedUserList
}

func (s *SignalingUseCase) DynamicSearch(userID string, geoLocation model.GeoLocation, searchDistance float64) []*model.UserInfo {
	userInfoList := s.userInfoRepository.FindAll()
	// DoSomething
	return userInfoList
}

func (s *SignalingUseCase) Send() {
}

func (s *SignalingUseCase) TwoPointsDistance(geoLocation model.GeoLocation, v model.GeoLocation) float64 {
	return 1000 * 6371 * math.Acos(math.Cos(geoLocation.Latitude*math.Pi/180)*math.Cos(v.Latitude*math.Pi/180)*math.Cos(v.Longitude*math.Pi/180-geoLocation.Longitude*math.Pi/180)+math.Sin(geoLocation.Latitude*math.Pi/180)*math.Sin(v.Latitude*math.Pi/180))
}
