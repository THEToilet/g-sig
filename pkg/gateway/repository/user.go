package repository

import (
	"g-sig/pkg/domain/model"
	"g-sig/pkg/domain/repository"
)

var _ repository.UserRepository = &UserRepository{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

type UserRepository struct {
}

func (u UserRepository) Find(id string) (*model.User, error) {
	panic("implement me")
}

func (u UserRepository) FindAll() ([]*model.User, error) {
	panic("implement me")
}

func (u UserRepository) Save(user model.User) error {
	panic("implement me")
}

func (u UserRepository) Delete(id string) error {
	panic("implement me")
}
