package model

// UserInfo ユーザの頻繁に変わる情報
type UserInfo struct {
	UserID      string  `json:"userID"`
	PublicIP    string  `json:"publicIP"`
	PublicPort  uint8   `json:"publicPort"`
	PrivateIP   string  `json:"privateIP"`
	PrivatePort uint8   `json:"privatePort"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

// GeoLocation ユーザの位置情報
type GeoLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Addr struct {
	IP   string `json:"ip"`
	Port uint8  `json:"port"`
}

// AlterUserInfo 別の案
type AlterUserInfo struct {
	UserID      string      `json:"userID"`
	PublicAddr  Addr        `json:"public"`
	PrivateAddr Addr        `json:"private"`
	GeoLocation GeoLocation `json:"geoLocation"`
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
