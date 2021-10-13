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

func (w *WSConnection) makeResponseMessage(code string, message string, actionType string) ([]byte, error) {
	status := respMessage.Status{
		Code:    code,
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

func (w *WSConnection) makeRegisterMessage(userID string) ([]byte, error) {
	registerResponse := respMessage.RegisterResponse{
		Status: w.makeStatusMessage("", "", ""),
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
	status := w.makeStatusMessage("", "", "")
	responseMessage, err := json.Marshal(status)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeDeleteMessage() ([]byte, error) {
	status := w.makeStatusMessage("", "", "")
	responseMessage, err := json.Marshal(status)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeSearchMessage(searchedUserList []*model.UserInfo) ([]byte, error) {
	searchResponse := respMessage.SearchResponse{
		Status:           w.makeStatusMessage("", "", ""),
		SearchedUserList: searchedUserList,
	}
	responseMessage, err := json.Marshal(searchResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil

}
func (w *WSConnection) makeSendMessage() ([]byte, error) {
	status := w.makeStatusMessage("", "", "")
	responseMessage, err := json.Marshal(status)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeStatusMessage(actionType string, code string, message string) respMessage.Status {
	status := respMessage.Status{
		Type:    actionType,
		Code:    code,
		Message: message,
	}
	return status
}

func (w *WSConnection) handleMessage(rawMessage []byte, pongTimer *time.Timer) {
	message := &model.Message{}
	if err := json.Unmarshal(rawMessage, &message); err != nil {
		w.logger.Fatal().Err(err)
	}
	if message == nil {
		w.logger.Fatal()
	}

	w.logger.Debug().Interface("message", message).Msg("Unmarshall message")

	switch message.Type {
	case "p2p":
		// pWebRTCのシグナリングサーバとして動く
	case "pong":
		stopTimer(pongTimer)
		pongTimer.Reset(time.Second * 10)
		w.logger.Info().Msg("pong")
	case "register":

		// userInfo登録
		registerMessage := &model.RegisterMessage{}
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

		w.receiveMessage <- responseMessage

	case "update":

		// userInfo更新
		updateMessage := &model.UpdateMessage{}
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

		w.receiveMessage <- responseMessage

	case "delete":

		// userInfo削除
		deleteMessage := &model.DeleteMessage{}
		if err := json.Unmarshal(rawMessage, &deleteMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg("delete")
		err := w.signalingUseCase.Delete(deleteMessage.UserInfo)
		if err != nil {
			return
		}

		responseMessage, err := w.makeDeleteMessage()
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.receiveMessage <- responseMessage

	case "search":

		// 周囲端末検索
		searchMessage := &model.SearchMessage{}
		if err := json.Unmarshal(rawMessage, &searchMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg("search")

		var searchedUserList []*model.UserInfo

		switch searchMessage.SearchType {
		case "static":
			searchedUserList = w.signalingUseCase.StaticSearch(searchMessage.UserInfo, searchMessage.SearchDistance)
		case "dynamic":
			searchedUserList = w.signalingUseCase.DynamicSearch(searchMessage.UserInfo, searchMessage.SearchDistance)
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
		w.receiveMessage <- responseMessage

	case "send":

		// 周囲に一斉送信
		sendMessage := &model.SendMessage{}
		if err := json.Unmarshal(rawMessage, &sendMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg("send")
		w.signalingUseCase.Send()

		responseMessage, err := w.makeSendMessage()
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.receiveMessage <- responseMessage

	default:
		w.logger.Debug().Interface("rawMessage", rawMessage).Msg("Invalid Message")

		responseMessage, err := w.makeResponseMessage("400", "Invalid Message", "undefined")
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.receiveMessage <- responseMessage
	}
}
