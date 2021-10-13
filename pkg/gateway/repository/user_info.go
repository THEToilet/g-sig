package repository

import (
	"g-sig/pkg/domain/model"
	"g-sig/pkg/domain/repository"
	"sync"
)

var _ repository.UserInfoRepository = &UserInfoRepository{}

var (
	//	UserInfoList =  map[string]model.UserInfo{}
	UserInfoList = &sync.Map{}
)

type UserInfoRepository struct {
}

func NewUserInfoRepository() *UserInfoRepository {
	return &UserInfoRepository{
	}
}

func (u UserInfoRepository) Find(userID string) (*model.UserInfo, error) {
	userInfo, ok := UserInfoList.Load(userID)
	if !ok {
		return nil, model.ErrUserNotFound
	}
	v, ok := userInfo.(model.UserInfo)
	if !ok {
		return nil, model.ErrUserNotFound
	}
	return &v, nil
}

func (u UserInfoRepository) FindAll() ([]*model.UserInfo, error) {
	var userInfoList []*model.UserInfo
	UserInfoList.Range(func(key, value interface{}) bool {
		v, ok := value.(model.UserInfo)
		if !ok {
			return false
		}
		userInfoList = append(userInfoList, &v)
		return true
	})
	return userInfoList, nil
}

func (u UserInfoRepository) Save(user model.UserInfo) error {
	_, ok := UserInfoList.Load(user.UserID)
	if ok {
		return model.ErrUserAlreadyExisted
	}
	UserInfoList.Store(user.UserID, user)
	return nil
}

func (u UserInfoRepository) Update(user model.UserInfo) error {
	_, ok := UserInfoList.Load(user.UserID)
	if !ok {
		return model.ErrUserNotFound
	}
	UserInfoList.Store(user.UserID, user)
	return nil
}

func (u UserInfoRepository) Delete(userID string) error {
	_, ok := UserInfoList.Load(userID)
	if !ok {
		return model.ErrUserNotFound
	}
	UserInfoList.Delete(userID)
	return nil
}
