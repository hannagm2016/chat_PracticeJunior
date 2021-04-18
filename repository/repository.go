package repository

import (
	//"fmt"
	"ChatPerson/models"
	"gorm.io/gorm"
)

type ChatModel struct {
	db *gorm.DB
}

var chats []models.Chats

type ChatModelImpl interface {
	FindContact(id float64) models.Contact
	FindContacts() []models.Contact
	FindChat(UserId float64) []models.Message
	FindChats() []models.Chats
	SaveMessage(mesage models.Messages)
}

func NewChatModel(db *gorm.DB) *ChatModel {
	return &ChatModel{
		db: db,
	}
}

func (p *ChatModel) FindChats() []models.Chats {

	Msgs := []models.Message{}
	Messages := []models.Message{}
	currentUser := 3                                                                                                       // тут будет ид текущего юзера
	p.db.Raw("SELECT id, user_from_id as user_id, text, time FROM messages where user_to_id = ?", currentUser).Scan(&Msgs) //пока возвращает первое сообщение
	for index, _ := range Msgs {
		Msgs[index].Type = "others"
	}
	Messages = Msgs
	p.db.Raw("SELECT id, user_to_id as user_id, text, time FROM messages where user_from_id = ?", currentUser).Scan(&Msgs) //пока возвращает первое сообщение
	for index, _ := range Msgs {
		Msgs[index].Type = "own"
		Messages = append(Messages, Msgs[index])
	}
	p.db.Raw("select con.* from (SELECT id, name FROM contacts where id !=?) as con join (SELECT DISTINCT  user_from_id user_id FROM messages  UNION SELECT DISTINCT user_to_id user_id from messages) as mes on con.id =mes.user_id", currentUser).Scan(&chats)
	//SELECT `name`, `time`, text` FROM `messages` GROUP BY `user_id` ORDER BY `id` DESC

	//select distinct `o`.`user` from (SELECT `id`, `user_from_id` `user` FROM `messages` where `user_from_id` !=3 UNION SELECT`id`,`user_to_id` `user` from messages where `user_to_id`!=3 order by `id` desc) `o`
	//   p.db.Raw("select con.* from (SELECT DISTINCT o.user from  (SELECT id, user_from_id user FROM messages where user_from_id !=? UNION SELECT id,user_to_id user from messages where user_to_id !=? order by id desc) o) mes  join  (SELECT id, name FROM contacts where id !=?) as con on mes.user=con.id", currentUser,currentUser,currentUser).Scan(&chats)

	for _, mess := range Messages {
		for index, chat := range chats {
			if chat.Id == mess.UserId {
				chats[index].Chat = append(chats[index].Chat, mess)
			}
		}
	}
	return chats
}
func (p *ChatModel) FindChat(UserId float64) []models.Message {
	messages := []models.Message{}
	p.db.Find(&messages, "user_id = ?", UserId)

	return messages
}

func (p *ChatModel) FindContacts() []models.Contact {
	contacts := []models.Contact{}
	p.db.Find(&contacts)
	return contacts
}
func (p *ChatModel) FindContact(UserId float64) models.Contact {
	contact := models.Contact{}
	p.db.Find(&contact, "id = ?", UserId)
	return contact
}

func (p *ChatModel) SaveMessage(message models.Messages) { //не инсертит
	p.db.Create(&message)
	/*	type message struct{
		    userToId, userFromId float64
		    text, time string
		}
			fmt.Println(&mes , "&&&&&")
		p.db.Create(&mes)*/
	//p.db.Exec("insert into messages (id, user_from_id, user_to_id, time, text) values(,3,?,?,?)", mess.UserId,mess.Time, mess.Text)

}
