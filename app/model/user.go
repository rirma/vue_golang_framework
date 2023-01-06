package model

//引っ張ってきたデータを割り当てるための構造体を用意
type User struct {
	Id       int
	Email    string
	Password string
	Name     string
}

func (User) TableName() string {
	return "users"
}