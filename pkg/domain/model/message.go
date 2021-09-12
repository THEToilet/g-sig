package model

type message struct {
	Type string `json:"type"`
}

type registerMessage struct {
	Type string `json:"type"`
	User user `json:"user"`
	UserInfo userInfo `json:"userInfo"`
}

type updateMessage struct {
	Type string `json:"type"`
	UserInfo userInfo `json:"userInfo"`
}

type searchMessage struct {
	Type string `json:"type"`
	UserInfo userInfo `json:"userInfo"`
	SearchType string `json:"searchType"`
	SearchDistance float64 `json:"searchDistance"`
}

type deleteMessage struct {
	Type string `json:"type"`
	UserInfo userInfo `json:"userInfo"`
}

type sendMessage struct {
	Type string `json:"type"`
	Message string `json:"message"`
}
