package handler

import (
	"encoding/json"
	"g-sig/pkg/domain/application"
	"g-sig/pkg/domain/model"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/zerolog"
	"net/http"
)

type signalingHandler struct {
	signalingUseCase *application.SignalingUseCase
	logger *zerolog.Logger
}

func NewSignalingHandler(useCase *application.SignalingUseCase, logger *zerolog.Logger) *signalingHandler {
	return &signalingHandler{
		signalingUseCase: useCase,
		logger: logger,
	}
}

func (h *signalingHandler)Signaling(writer http.ResponseWriter, request *http.Request){
	conn, _, _, err := ws.UpgradeHTTP(request, writer)
	if err != nil{
		h.logger.Fatal().Err(err)
		return
	}
	h.logger.Info().Msg("Connection Start")
	go func() {
		defer conn.Close()
		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				return
			}
			message := &model.Message{}
			err := json.Unmarshal(msg, &message);
			if err != nil {
				h.logger.Fatal().Err(err)
			}
			h.logger.Info().Msg(string(msg))
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				return
			}
		}
	}()
}
