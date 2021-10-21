package handler

import (
	"encoding/json"
	"g-sig/pkg/domain/model"
	respMessage "g-sig/pkg/server/message"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
	"time"
)

// 今はエラー用だけに使われている
func (w *WSConnection) makeResponseMessage(message string, actionType string) ([]byte, error) {
	status := respMessage.Response{
		Message: message,
		Type:    actionType,
	}
	responseMessage, err := json.Marshal(status)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}

func (w *WSConnection) sendMessage(message []byte) error {
	// TODO opCode要検討
	w.logger.Debug().Msg(string(message))
	if err := wsutil.WriteServerMessage(w.conn, ws.OpText, message); err != nil {
		w.logger.Fatal().Err(err)
		return err
	}
	return nil
}

func (w *WSConnection) makePingMessage() ([]byte, error) {
	pingRequest := respMessage.PingRequest{
		Type: "ping",
	}
	requestMessage, err := json.Marshal(pingRequest)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return requestMessage, nil
}

func (w *WSConnection) makeRegisterMessage(userID string) ([]byte, error) {
	registerResponse := respMessage.RegisterResponse{
		Type:   "register",
		UserID: userID,
	}
	responseMessage, err := json.Marshal(registerResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeUpdateMessage() ([]byte, error) {
	updateResponse := respMessage.UpdateResponse{
		Type:    "update",
		Message: "",
	}
	responseMessage, err := json.Marshal(updateResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeDeleteMessage() ([]byte, error) {
	deleteResponse := respMessage.DeleteResponse{
		Type:    "delete",
		Message: "",
	}
	responseMessage, err := json.Marshal(deleteResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeSearchMessage(searchedUserList []*model.UserInfo) ([]byte, error) {
	searchResponse := respMessage.SearchResponse{
		Type:                "search",
		Message:             "",
		SurroundingUserList: searchedUserList,
	}
	responseMessage, err := json.Marshal(searchResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil

}
func (w *WSConnection) makeSendMessage() ([]byte, error) {
	sendResponse := respMessage.SendResponse{
		Type:    "send",
		Message: "",
	}
	responseMessage, err := json.Marshal(sendResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}

func (w *WSConnection) handleMessage(rawMessage []byte, pongTimer *time.Timer) {
	message := &respMessage.JudgeMessageType{}
	if err := json.Unmarshal(rawMessage, &message); err != nil {
		w.logger.Fatal().Err(err)
	}
	if message == nil {
		w.logger.Fatal()
	}

	switch message.Type {
	case "p2p":
		// pWebRTCのシグナリングサーバとして動く
	case "pong":
		stopTimer(pongTimer)
		pongTimer.Reset(time.Second * 10)
		w.logger.Info().Msg("pong")
	case "register":

		// userInfo登録
		registerMessage := &respMessage.RegisterRequest{}
		if err := json.Unmarshal(rawMessage, &registerMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg("register")

		userID, err := uuid.NewUUID()
		if err != nil {
			return
		}

		err = w.signalingUseCase.Register(userID.String(), registerMessage.GeoLocation)
		if err != nil {
			w.logger.Fatal().Err(err)
		}

		w.isRegistered = true
		w.userID = userID.String()

		responseMessage, err := w.makeRegisterMessage(userID.String())

		w.sendingMessage <- responseMessage

	case "update":

		// userInfo更新
		updateMessage := &respMessage.UpdateRequest{}
		if err := json.Unmarshal(rawMessage, &updateMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg("update")
		err := w.signalingUseCase.Update(updateMessage.UserInfo)
		if err != nil {
			return
		}

		responseMessage, err := w.makeUpdateMessage()
		if err != nil {
			w.logger.Fatal().Err(err)
		}

		w.sendingMessage <- responseMessage

	case "delete":

		// userInfo削除
		deleteMessage := &respMessage.DeleteRequest{}
		if err := json.Unmarshal(rawMessage, &deleteMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg("delete")
		err := w.signalingUseCase.Delete(w.userID)
		if err != nil {
			return
		}

		responseMessage, err := w.makeDeleteMessage()
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.sendingMessage <- responseMessage

	case "search":

		// TODO: ユーザを検索した際に誰も該当者がいないときの動作をもう少し考える
		// 周囲端末検索
		searchMessage := &respMessage.SearchRequest{}
		if err := json.Unmarshal(rawMessage, &searchMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg("search")

		var searchedUserList []*model.UserInfo

		switch searchMessage.SearchType {
		case "static":
			searchedUserList = w.signalingUseCase.StaticSearch(w.userID, searchMessage.GeoLocation, searchMessage.SearchDistance)
		case "dynamic":
			searchedUserList = w.signalingUseCase.DynamicSearch(w.userID, searchMessage.GeoLocation, searchMessage.SearchDistance)
		default:
			w.logger.Info().Msg("invalid type")
		}

		if searchedUserList == nil {
			searchedUserList = append(searchedUserList, &model.UserInfo{})
		}

		responseMessage, err := w.makeSearchMessage(searchedUserList)
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.sendingMessage <- responseMessage

	case "send":

		// 周囲に一斉送信
		sendMessage := &respMessage.SendRequest{}
		if err := json.Unmarshal(rawMessage, &sendMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg("send")
		w.signalingUseCase.Send()

		responseMessage, err := w.makeSendMessage()
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.sendingMessage <- responseMessage

	default:
		w.logger.Debug().Interface("rawMessage", rawMessage).Msg("Invalid RequestType")

		responseMessage, err := w.makeResponseMessage("Invalid RequestType", "undefined")
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.sendingMessage <- responseMessage
	}
}
