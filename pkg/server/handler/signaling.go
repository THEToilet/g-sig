package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"g-sig/pkg/domain/application"
	"g-sig/pkg/domain/model"
	response_message "g-sig/pkg/server/message"
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

	// コンテキストを使う
	// タイムアウト処理とかで使える
	// 今までアクセス時の情報が入ると思ってたけどそういうわけではないらしい
	// Key：Valueのデータ構造も持つ
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	receiveMessage := make(chan []byte, 100)
	sendMessage := make(chan []byte, 100)
	wsConnection := WSConnection{
		conn:   conn,
		receiveMessage: receiveMessage,
		sendMessage: sendMessage,
		logger: h.logger,
	}
	go wsConnection.receiver(ctx)
	go wsConnection.selector(ctx)
	// TODO: うつす
	defer cancel()
}

func (h *signalingHandler) makeResponseMessage(code string, message string, actionType string) ([]byte, error) {
	status := response_message.Status{
		Code:    code,
		Message: message,
		Type:    actionType,
	}
	responseMessage, err := json.Marshal(status)
	if err != nil {
		h.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}

func (h *signalingHandler) sendMessage(message []byte) error {
	if err := wsutil.WriteServerMessage(conn, op, message); err != nil {
		h.logger.Fatal().Err(err)
		return err
	}
	return nil
}

func (h *signalingHandler) handleMessage() {
	var responseMessage []byte
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

		userID, err := h.signalingUseCase.Register(registerMessage.UserInfo)

		status := response_message.Status{
			Code:    "200",
			Message: "OK",
			Type:    "register",
		}
		registerResponse := response_message.RegisterResponse{
			Status: status,
			UserID: userID,
		}

		responseMessage, err = json.Marshal(registerResponse)
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
		err := h.signalingUseCase.Update(updateMessage.UserInfo)
		if err != nil {
			return
		}

		status := response_message.Status{
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
		err := h.signalingUseCase.Delete(deleteMessage.UserInfo)
		if err != nil {
			return
		}

		status := response_message.Status{
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

		if searchedUserList == nil {
			searchedUserList = append(searchedUserList, &model.UserInfo{})
		}

		// Debug
		searchedUserList = append(searchedUserList, &model.UserInfo{
			UserID:      "test1",
			PublicIP:    "",
			PublicPort:  0,
			PrivateIP:   "",
			PrivatePort: 0,
			Latitude:    35.943207099999995,
			Longitude:   139.6211672,
		})
		searchedUserList = append(searchedUserList, &model.UserInfo{
			UserID:      "test2",
			PublicIP:    "",
			PublicPort:  0,
			PrivateIP:   "",
			PrivatePort: 0,
			Latitude:    35.948658,
			Longitude:   139.640718,
		})
		searchedUserList = append(searchedUserList, &model.UserInfo{
			UserID:      "test3",
			PublicIP:    "",
			PublicPort:  0,
			PrivateIP:   "",
			PrivatePort: 0,
			Latitude:    35.950828,
			Longitude:   139.651533,
		})

		status := response_message.Status{
			Code:    "200",
			Message: "OK",
			Type:    "search",
		}
		searchResponse := response_message.SearchResponse{
			Status:           status,
			SearchedUserList: searchedUserList,
		}
		responseMessage, err = json.Marshal(searchResponse)
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

		responseMessage, err = h.makeResponseMessage("200", "OK", "send")
		if err != nil {
			h.logger.Fatal().Err(err)
		}

	default:
		h.logger.Info().Msg("invalid message")

		responseMessage, err = h.makeResponseMessage("400", "Invalid Message", "undefined")
		if err != nil {
			h.logger.Fatal().Err(err)
		}
	}

}
