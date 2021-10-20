package message

import "g-sig/pkg/domain/model"

type PingRequest struct {
	Type string `json:"type"`
}

type RegisterRequest struct {
	Type        string            `json:"type"`
	GeoLocation model.GeoLocation `json:"geoLocation"`
}

type UpdateRequest struct {
	Type     string         `json:"type"`
	UserInfo model.UserInfo `json:"userInfo"`
}

type SearchRequest struct {
	Type           string            `json:"type"`
	SearchType     string            `json:"searchType"`
	GeoLocation    model.GeoLocation `json:"geoLocation"`
	SearchDistance float64           `json:"searchDistance"`
}

type DeleteRequest struct {
	Type string `json:"type"`
}

type SendRequest struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
