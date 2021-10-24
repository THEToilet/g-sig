package handler

import (
	"errors"
	"g-sig/pkg/domain/application"
	"g-sig/pkg/domain/model"
	"github.com/rs/zerolog"
	"net"
	"sync"
)

var (
	connections = &sync.Map{}
)

type Connections struct {
	signalingUseCase *application.SignalingUseCase
	logger           *zerolog.Logger
}

func NewConnections(logger *zerolog.Logger) *Connections {
	return &Connections{
		logger: logger,
	}
}

func (c *Connections) Save(userID string, conn *net.Conn) error {
	_, ok := connections.Load(userID)
	if ok {
		return model.ErrUserAlreadyExisted
	}
	connections.Store(userID, conn)
	return nil
}

func (c *Connections) Delete(userID string) error {
	_, ok := connections.Load(userID)
	if !ok {
		return model.ErrUserNotFound
	}
	connections.Delete(userID)
	return nil
}

func (c *Connections) Find(userID string) (*net.Conn, error) {
	conn, ok := connections.Load(userID)
	if !ok {
		return nil, errors.New("d")
	}
	v, ok := conn.(*net.Conn)
	if !ok {
		return nil, model.ErrUserNotFound
	}
	return v, nil
}
