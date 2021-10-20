package model

// UserInfo ユーザの頻繁に変わる情報
type UserInfo struct {
	UserID      string      `json:"userID"`
	GeoLocation GeoLocation `json:"geoLocation"`
}

// GeoLocation ユーザの位置情報
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// NewUserInfo 新しいUserInfoを生成したポインタを返す
func NewUserInfo(userID string, latitude float64, longitude float64) *UserInfo {
	return &UserInfo{
		UserID: userID,
		GeoLocation: GeoLocation{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}
}
