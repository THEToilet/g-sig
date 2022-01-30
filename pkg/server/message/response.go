package message

import "g-sig/pkg/domain/model"

type JudgeMessageType struct {
	Type string `json:"type"`
}

type RegisterResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	UserID  string `json:"userID"`
}
type UpdateResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type SearchResponse struct {
	Type                string            `json:"type"`
	Message             string            `json:"message"`
	SurroundingUserList []*model.UserInfo `json:"surroundingUserList"`
	ResponseID          string            `json:"responseID"`
}

type DeleteResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type SendResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Response struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
