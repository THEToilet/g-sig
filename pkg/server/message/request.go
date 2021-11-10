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

type OfferMessage struct {
	Type        string `json:"type"`
	SDP         string `json:"sdp"`
	Destination string `json:"destination"`
}

type AnswerMessage struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	SDP         string `json:"sdp"`
}

type ICEMessage struct {
	Type        string `json:"type"`
	ICE         string `json:"ice"`
}

type CloseMessage struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
}
