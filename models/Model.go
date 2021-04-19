package models

import (
//"gorm.io/gorm"
)

type Contact struct {
	Id    float64
	Name  string
	Phone string
	//Password    string
	Status string
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
	UserToId   float64
	UserFromId float64
	Time       string
	Text       string
}
