package repository

import (
	"errors"
	"g-sig/pkg/domain/model"
	"g-sig/pkg/domain/repository"
	"github.com/rs/zerolog"
)

var _ repository.UserInfoRepository = &userInfoRepository{}

var (
	UserInfoList =  map[string]model.UserInfo{}
)

type userInfoRepository struct {
	logger *zerolog.Logger
}

func NewUserInfoRepository(logger *zerolog.Logger) *userRepository {
	return &userRepository{
		logger: logger,
	}
}

func (u userInfoRepository) Find(userID string) (*model.UserInfo, error) {
	userInfo, ok := UserInfoList[userID]
	if !ok {
		return nil, errors.New("user not found")
	}
	return &userInfo, nil
}

func (u userInfoRepository) FindAll() ([]*model.UserInfo, error) {
	var userInfoList []*model.UserInfo
	for _, userInfo := range UserInfoList {
		userInfoList = append(userInfoList, &userInfo)
	}
	return userInfoList, nil
}

func (u userInfoRepository) Save(user model.UserInfo) error {
	_, ok := UserInfoList[user.UserID]
	if ok {
		return  errors.New("user found")
	}
	UserInfoList[user.UserID] = user
	return nil
}

func (u userInfoRepository) Update(user model.UserInfo) error {
	_, ok := UserInfoList[user.UserID]
	if !ok {
		return errors.New("user not found")
	}
	UserInfoList[user.UserID] = user
	return nil
}

func (u userInfoRepository) Delete(userID string) error {
	_, ok := UserInfoList[userID]
	if !ok {
		return errors.New("user not found")
	}
	delete(UserInfoList, userID)
	return nil
}

