package model

// user 永続化するユーザ情報
type user struct {
	UserID string
	UserName string
}

// NewUser 新しいユーザ情報を生成してポインタを返す
func NewUser(userID string, userName string) *user {
	return &user{
		UserID:  userID,
		UserName: userName,
	}
}
