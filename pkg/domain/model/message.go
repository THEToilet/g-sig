package model

type Message struct {
	Type string `json:"type"`
}

type RegisterMessage struct {
	Type     string   `json:"type"`
	UserInfo UserInfo `json:"userInfo"`
}

type UpdateMessage struct {
	Type     string   `json:"type"`
	UserInfo UserInfo `json:"userInfo"`
}

type SearchMessage struct {
	Type           string   `json:"type"`
	UserInfo       UserInfo `json:"userInfo"`
	SearchType     string   `json:"searchType"`
	SearchDistance float64  `json:"searchDistance"`
}

type DeleteMessage struct {
	Type     string   `json:"type"`
	UserInfo UserInfo `json:"userInfo"`
}

type SendMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
