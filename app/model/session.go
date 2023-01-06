package model

//引っ張ってきたデータを割り当てるための構造体を用意
type Session struct {
	Id       int `gorm:"primary_key"`
	UserId   int
	Payload  string
	LastActivity int64
}

func (Session) TableName() string {
	return "sessions"
}