package model

// User 永続化するユーザ情報
type User struct {
	UserID string
	UserName string
}

// NewUser 新しいユーザ情報を生成してポインタを返す
func NewUser(userID string, userName string) *User {
	return &User{
		UserID:  userID,
		UserName: userName,
	}
}
