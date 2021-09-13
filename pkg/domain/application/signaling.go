package application

import "github.com/rs/zerolog"

type SignalingUseCase struct {
	logger *zerolog.Logger
}

func NewSignalingUseCase(logger *zerolog.Logger) *SignalingUseCase{
	return &SignalingUseCase{
		logger: logger,
	}
}