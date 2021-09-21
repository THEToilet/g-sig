package repository

import (
	"g-sig/pkg/domain/model"
	"g-sig/pkg/domain/repository"
	"github.com/rs/zerolog"
)

var _ repository.UserRepository = &userRepository{}

type userRepository struct {
	logger *zerolog.Logger
}

func (u userRepository) Find(id string) (*model.User, error) {
	panic("implement me")
}

func (u userRepository) FindAll() ([]*model.User, error) {
	panic("implement me")
}

func (u userRepository) Save(user model.User) error {
	panic("implement me")
}

func (u userRepository) Delete(id string) error {
	panic("implement me")
}

func NewUserRepository(logger *zerolog.Logger) *userRepository {
	return &userRepository{
		logger: logger,
	}
}
