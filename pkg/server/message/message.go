package message

import "g-sig/pkg/domain/model"

type Status struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type SearchResponse struct {
	Status Status `json:"status"`
	SearchedUserList []*model.UserInfo `json:"searchedUserList"`
}