package handler

import (
	"encoding/json"
	"g-sig/pkg/domain/application"
	"g-sig/pkg/domain/model"
	message2 "g-sig/pkg/server/message"
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
		var responseMessage []byte
		defer conn.Close()
		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				h.logger.Fatal().Err(err)
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

				status := message2.Status{
					Code:    "200",
					Message: "OK",
					Type:    "register",
				}
				responseMessage, err = json.Marshal(status)
				if err != nil {
					h.logger.Fatal().Err(err)
				}

			case "update":

				// userInfo更新
				updateMessage := &model.UpdateMessage{}
				if err := json.Unmarshal(msg, &updateMessage); err != nil {
					h.logger.Fatal().Err(err)
				}
				h.logger.Info().Msg("update")
				h.signalingUseCase.Update(updateMessage.UserInfo)

				status := message2.Status{
					Code:    "200",
					Message: "OK",
					Type:    "update",
				}
				responseMessage, err = json.Marshal(status)
				if err != nil {
					h.logger.Fatal().Err(err)
				}

			case "delete":

				// userInfo削除
				deleteMessage := &model.DeleteMessage{}
				if err := json.Unmarshal(msg, &deleteMessage); err != nil {
					h.logger.Fatal().Err(err)
				}
				h.logger.Info().Msg("delete")
				h.signalingUseCase.Delete(deleteMessage.UserInfo)

				status := message2.Status{
					Code:    "200",
					Message: "OK",
					Type:    "delete",
				}
				responseMessage, err = json.Marshal(status)
				if err != nil {
					h.logger.Fatal().Err(err)
				}

			case "search":

				// 周囲端末検索
				searchMessage := &model.SearchMessage{}
				if err := json.Unmarshal(msg, &searchMessage); err != nil {
					h.logger.Fatal().Err(err)
				}
				h.logger.Info().Msg("search")

				var searchedUserList []*model.UserInfo

				switch searchMessage.SearchType {
				case "static":
					searchedUserList, err = h.signalingUseCase.StaticSearch(searchMessage.UserInfo, searchMessage.SearchDistance)
				case "dynamic":
					searchedUserList, err = h.signalingUseCase.DynamicSearch(searchMessage.UserInfo, searchMessage.SearchDistance)
				default:
					h.logger.Info().Msg("invalid type")
				}

				status := message2.Status{
					Code:    "200",
					Message: "OK",
					Type:    "search",
				}
				tmp := message2.SearchResponse{
					Status:           status,
					SearchedUserList: searchedUserList,
				}
				responseMessage, err = json.Marshal(tmp)
				if err != nil {
					h.logger.Fatal().Err(err)
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
			err = wsutil.WriteServerMessage(conn, op, responseMessage)
			if err != nil {
				h.logger.Fatal().Err(err)
			}
		}
	}()
}
