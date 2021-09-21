package model

// userInfo ユーザの頻繁に変わる情報
type UserInfo struct {
	UserID      string
	PublicIP    string
	PublicPort  uint8
	PrivateIP   string
	PrivatePort uint8
	Latitude    float64
	Longitude   float64
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
