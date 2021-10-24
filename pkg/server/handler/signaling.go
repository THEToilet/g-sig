package handler

import (
	"context"
	"g-sig/pkg/domain/application"
	"github.com/gobwas/ws"
	"github.com/rs/zerolog"
	"net/http"
)

type signalingHandler struct {
	signalingUseCase *application.SignalingUseCase
	connections      *Connections
	logger           *zerolog.Logger
}

func NewSignalingHandler(useCase *application.SignalingUseCase, connections *Connections, logger *zerolog.Logger) *signalingHandler {
	return &signalingHandler{
		signalingUseCase: useCase,
		connections:      connections,
		logger:           logger,
	}
}

func (h *signalingHandler) Signaling(writer http.ResponseWriter, request *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(request, writer)
	if err != nil {
		h.logger.Fatal().Err(err)
		return
	}
	h.logger.Info().Msg("Connection Start")
	//h.logger.Debug().Caller().Msg("")

	// NOTE: goroutineのキャンセル処理に使う
	ctx, cancel := context.WithCancel(context.Background())

	receiveMessage := make(chan []byte, 100)
	sendingMessage := make(chan []byte, 100)
	wsConnection := NewWSConnection(conn, receiveMessage, sendingMessage, h.signalingUseCase, h.connections, h.logger)
	go wsConnection.selector(ctx, cancel)
	go wsConnection.receiver(ctx)
}
