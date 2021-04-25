package models

import (
//"gorm.io/gorm"
)

type Contact struct {
	Id    float64
	Name  string
	Phone string
	Password    []byte
	Email string
	Status string
	Relation string
}
type Chats struct {
	Id float64
	MesId float64
	//UserId float64
	Name string
	Time string
	Text string
	Chat []Message `gorm:"foreignKey:UserId;references:Id"` //`gorm:"-"`
}
type Message struct {
	Id     float64
	UserId float64
	Time   string
	Text   string
	Type   string
}
type Messages struct {
    Id float64
	UserToId   float64
	UserFromId float64
	Time       string
	Text       string
}

type BaseModel struct {

	Token string `json:"token"`
	User Contact `json:"user"`
}
type Relations struct{
    Relation string
    UserId, UserTo float64
}

