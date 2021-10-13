package message

import "g-sig/pkg/domain/model"

type Status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type SearchResponse struct {
	Status           Status            `json:"status"`
	SearchedUserList []*model.UserInfo `json:"searchedUserList"`
}
type RegisterResponse struct {
	Status Status `json:"status"`
	UserID string `json:"userID"`
}
type AlterRegisterResponse struct {
	Type   string `json:"type"`
	UserID string `json:"userID"`
}
