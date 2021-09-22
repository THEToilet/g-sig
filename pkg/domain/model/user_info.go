package model

// UserInfo ユーザの頻繁に変わる情報
type UserInfo struct {
	UserID      string   `json:"userID"`
	PublicIP    string   `json:"publicIP"`
	PublicPort  uint8    `json:"publicPort"`
	PrivateIP   string   `json:"privateIP"`
	PrivatePort uint8    `json:"privatePort"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
}

// NewUserInfo 新しいUserInfoを生成したポインタを返す
func NewUserInfo(userID string, publicIP string, publicPort uint8, privateIP string, privatePort uint8, latitude float64, longitude float64) *UserInfo {
	return &UserInfo{
		UserID:      userID,
		PublicIP:    publicIP,
		PublicPort:  publicPort,
		PrivateIP:   privateIP,
		PrivatePort: privatePort,
		Latitude:    latitude,
		Longitude:   longitude,
	}
}
