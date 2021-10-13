package model

import "errors"

var (
	ErrHoge = errors.New("hoge")
	ErrUserNotFound = errors.New("user not found")
	ErrUserAlreadyExisted = errors.New("user already existed")
)
