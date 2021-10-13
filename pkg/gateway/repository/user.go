package repository

import (
	"g-sig/pkg/domain/model"
	"g-sig/pkg/domain/repository"
	"github.com/rs/zerolog"
)

var _ repository.UserRepository = &UserRepository{}

type UserRepository struct {
	logger *zerolog.Logger
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

func NewUserRepository(logger *zerolog.Logger) *UserRepository {
	return &UserRepository{
		logger: logger,
	}
}
