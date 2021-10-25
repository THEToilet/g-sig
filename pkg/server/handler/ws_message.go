package handler

import (
	"encoding/json"
	"g-sig/pkg/domain/model"
	mess "g-sig/pkg/server/message"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
	"net"
	"time"
)

// 今はエラー用だけに使われている
func (w *WSConnection) makeResponseMessage(message string, actionType string) ([]byte, error) {
	status := mess.Response{
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

func (w *WSConnection) sendMessage(conn *net.Conn, message []byte) error {
	// TODO opCode要検討
	w.logger.Debug().Msg(string(message))
	if err := wsutil.WriteServerMessage(*conn, ws.OpText, message); err != nil {
		w.logger.Fatal().Err(err)
		return err
	}
	return nil
}

func (w *WSConnection) makePingMessage() ([]byte, error) {
	pingRequest := mess.PingRequest{
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
	registerResponse := mess.RegisterResponse{
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
	updateResponse := mess.UpdateResponse{
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
	deleteResponse := mess.DeleteResponse{
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
	searchResponse := mess.SearchResponse{
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
	sendResponse := mess.SendResponse{
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
func (w *WSConnection) makeOfferMessage(sdp string, destination string) ([]byte, error) {
	offerResponse := mess.OfferMessage{
		Type:        "offer",
		SDP:         sdp,
		Destination: destination,
	}
	responseMessage, err := json.Marshal(offerResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeAnswerMessage(sdp string, destination string) ([]byte, error) {
	answerResponse := mess.AnswerMessage{
		Type:        "answer",
		SDP:         sdp,
		Destination: destination,
	}
	responseMessage, err := json.Marshal(answerResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeCloseMessage(destination string) ([]byte, error) {
	closeResponse := mess.CloseMessage{
		Type:        "close",
		Destination: destination,
	}
	responseMessage, err := json.Marshal(closeResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}
func (w *WSConnection) makeIceMessage(ice string) ([]byte, error) {
	iceResponse := mess.ICEMessage{
		Type:        "ice",
		ICE:         ice,
	}
	responseMessage, err := json.Marshal(iceResponse)
	if err != nil {
		w.logger.Fatal().Err(err)
		return nil, err
	}
	return responseMessage, nil
}

func (w *WSConnection) handleMessage(rawMessage []byte, pongTimer *time.Timer) {
	message := &mess.JudgeMessageType{}
	if err := json.Unmarshal(rawMessage, &message); err != nil {
		w.logger.Fatal().Err(err)
	}
	if message == nil {
		w.logger.Fatal()
	}

	switch message.Type {
	case "pong":
		stopTimer(pongTimer)
		pongTimer.Reset(time.Second * 10)
		w.logger.Info().Msg("pong")
	case "register":

		// userInfo登録
		registerMessage := &mess.RegisterRequest{}
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

		w.connections.Save(w.userID, &w.conn)

		responseMessage, err := w.makeRegisterMessage(userID.String())

		w.sendingMessage <- responseMessage

	case "update":

		// userInfo更新
		updateMessage := &mess.UpdateRequest{}
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
		deleteMessage := &mess.DeleteRequest{}
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
		searchMessage := &mess.SearchRequest{}
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
		sendMessage := &mess.SendRequest{}
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

	case "offer":
		offerMessage := &mess.OfferMessage{}
		w.logger.Info().Msg(offerMessage.Type)
		if err := json.Unmarshal(rawMessage, &offerMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		destinationConn, err := w.connections.Find(offerMessage.Destination)
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Interface("destinationConn", destinationConn)
		if destinationConn == nil {
			w.logger.Info().Msg("destination is nil")
			return
		}

		// NOTE: 相手のユーザIDを保存
		w.destination = offerMessage.Destination

		// NOTE: ここでユーザIDを交換
		responseMessage, err := w.makeOfferMessage(offerMessage.SDP, w.userID)
		if err != nil {
			w.logger.Fatal().Err(err)
		}

		if err = w.sendMessage(destinationConn, responseMessage); err != nil {
			w.logger.Fatal().Err(err)
		}

	case "answer":
		answerMessage := &mess.AnswerMessage{}
		w.logger.Info().Msg(answerMessage.Type)
		if err := json.Unmarshal(rawMessage, &answerMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		destinationConn, err := w.connections.Find(answerMessage.Destination)
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		if destinationConn == nil {
			w.logger.Info().Msg("destination is nil")
			return
		}

		// NOTE: 相手のユーザIDを保存
		w.destination = answerMessage.Destination

		// NOTE: ここでユーザIDを交換
		responseMessage, err := w.makeAnswerMessage(answerMessage.SDP, w.userID)
		if err != nil {
			w.logger.Fatal().Err(err)
		}

		if err = w.sendMessage(destinationConn, responseMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
	case "ice":
		iceMessage := &mess.ICEMessage{}
		w.logger.Info().Msg(iceMessage.Type)
		if err := json.Unmarshal(rawMessage, &iceMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		w.logger.Info().Msg(w.destination)
		destinationConn, err := w.connections.Find(w.destination)
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		if destinationConn == nil {
			w.logger.Info().Msg("destination is nil")
			return
		}

		// NOTE: ここでユーザIDを交換
		responseMessage, err := w.makeIceMessage(iceMessage.ICE)
		if err != nil {
			w.logger.Fatal().Err(err)
		}

		if err = w.sendMessage(destinationConn, responseMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
	case "close":
		closeMessage := &mess.CloseMessage{}
		w.logger.Info().Msg(closeMessage.Type)
		if err := json.Unmarshal(rawMessage, &closeMessage); err != nil {
			w.logger.Fatal().Err(err)
		}
		destinationConn, err := w.connections.Find(closeMessage.Destination)
		if err != nil {
			w.logger.Fatal().Err(err)
		}

		// NOTE: ここでユーザIDを交換
		responseMessage, err := w.makeCloseMessage(w.userID)
		if err != nil {
			w.logger.Fatal().Err(err)
		}

		if err = w.sendMessage(destinationConn, responseMessage); err != nil {
			w.logger.Fatal().Err(err)
		}

	default:
		w.logger.Debug().Interface("rawMessage", rawMessage).Msg("Invalid RequestType")

		responseMessage, err := w.makeResponseMessage("Invalid RequestType", "undefined")
		if err != nil {
			w.logger.Fatal().Err(err)
		}
		w.sendingMessage <- responseMessage
	}
}
