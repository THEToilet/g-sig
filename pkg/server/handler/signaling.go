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
	logger           *zerolog.Logger
}

func NewSignalingHandler(useCase *application.SignalingUseCase, logger *zerolog.Logger) *signalingHandler {
	return &signalingHandler{
		signalingUseCase: useCase,
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
	go func() {
		defer conn.Close()
		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				return
			}
			h.logger.Info().Msg(string(msg))

			message := &model.Message{}
			if err := json.Unmarshal(msg, &message); err != nil {
				h.logger.Fatal().Err(err)
			}
			if message == nil {
				h.logger.Fatal()
			}

			switch message.Type {
			case "register":

				// userInfo登録
				registerMessage := &model.RegisterMessage{}
				if err := json.Unmarshal(msg, &registerMessage); err != nil {
					h.logger.Fatal().Err(err)
				}
				h.logger.Info().Msg("register")
				h.signalingUseCase.Register(registerMessage.UserInfo)

			case "update":

				// userInfo更新
				updateMessage := &model.UpdateMessage{}
				if err := json.Unmarshal(msg, &updateMessage); err != nil {
					h.logger.Fatal().Err(err)
				}
				h.logger.Info().Msg("update")
				h.signalingUseCase.Update(updateMessage.UserInfo)

			case "delete":

				// userInfo削除
				deleteMessage := &model.DeleteMessage{}
				if err := json.Unmarshal(msg, &deleteMessage); err != nil {
					h.logger.Fatal().Err(err)
				}
				h.logger.Info().Msg("delete")
				h.signalingUseCase.Delete(deleteMessage.UserInfo)

			case "search":

				// 周囲端末検索
				searchMessage := &model.SearchMessage{}
				if err := json.Unmarshal(msg, &searchMessage); err != nil {
					h.logger.Fatal().Err(err)
				}
				h.logger.Info().Msg("search")

				switch searchMessage.SearchType {
				case "static":
					h.signalingUseCase.StaticSearch(searchMessage.UserInfo, searchMessage.SearchDistance)
				case "dynamic":
					h.signalingUseCase.DynamicSearch(searchMessage.UserInfo, searchMessage.SearchDistance)
				default:
					h.logger.Info().Msg("invalid type")
				}

			case "send":

				// 周囲に一斉送信
				sendMessage := &model.SendMessage{}
				if err := json.Unmarshal(msg, &sendMessage); err != nil {
					h.logger.Fatal().Err(err)
				}
				h.logger.Info().Msg("send")
				h.signalingUseCase.Send()

			default:
				h.logger.Info().Msg("invalid message")
			}

			// ここでステータスコードを返す?
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				h.logger.Fatal().Err(err)
			}
		}
	}()
}
