package repository

import "g-sig/pkg/domain/model"

type UserInfoRepository interface {
	Find(id string) (*model.UserInfo, error)
	FindAll() ([]*model.UserInfo, error)
	Save(user model.UserInfo) error
	Update(user model.UserInfo) error
	Delete(userID string) error
}
