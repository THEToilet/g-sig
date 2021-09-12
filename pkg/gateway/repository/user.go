package repository

import (
	"g-sig/pkg/domain/repository"
)

var _ repository.UserRepository = &userRepository{}

type userRepository struct {

}

func NewUserRepository() *userRepository {
	return &userRepository{

	}
}