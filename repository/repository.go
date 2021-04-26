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
	FindContact(userId float64) models.Contact
	UpdateMyContact(contact models.Contact, userId float64) models.Contact
	FindContacts(userId float64) []models.Contact
	FindChat(UserId float64) []models.Message
	FindChats(Uid float64) []models.Chats
	SaveMessage(mesage models.Messages)
	    CreateUser(contact models.Contact)
    	FindCustomerByName(userName string) models.Contact
    	FindCustomerById(userId int) models.Contact
    	FindCustomerByEmail(email string) models.Contact
    	FindCustomerId(name string) float64
    	SetRelation(relation models.Relations)
    	DeleteRelation(relation models.Relations)
    	ChangeRelation(relation models.Relations)
}

func NewChatModel(db *gorm.DB) *ChatModel {
	return &ChatModel{
		db: db,
	}
}
func (p *ChatModel) FindChats(currentUser float64) []models.Chats {
Msgs := []models.Messages{}
	Messages := []models.Message{}
	message := models.Message{}
	p.db.Raw("SELECT * FROM messages WHERE user_from_id = ? or user_to_id = ? order by id asc", currentUser, currentUser).Scan(&Msgs)
	for _, msg := range Msgs {
	if msg.UserFromId == currentUser {
	    message.Id=msg.Id
	    message.UserId= msg.UserToId
	    message.Time= msg.Time
	    message.Text= msg.Text
	    message.Type= "own"} else if msg.UserToId == currentUser {
   	    message.Id=msg.Id
        message.UserId= msg.UserFromId
        message.Time= msg.Time
        message.Text= msg.Text
        message.Type= "other"}
        Messages = append(Messages, message)
	}

	p.db.Raw("select mes.*, con.name from (select max(id) mes_id, o.user id, text, time from (SELECT id, user_from_id user, text, time FROM messages where user_to_id=? union SELECT id, user_to_id user, text, time FROM messages where user_from_id =?) o group by o.user) mes join (SELECT id, name FROM contacts where id not in (select user_to from relations where user_id=? and relation = 'Blocked')) as con on mes.id=con.id order by mes_id desc", currentUser,currentUser,currentUser).Scan(&chats)


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
func (p *ChatModel) FindContacts(userId float64) []models.Contact {
	contacts := []models.Contact{}
	p.db.Raw("select con.*, rel.relation from (SELECT id, name, phone, email, status FROM `contacts` where id !=?) con left join (select relation, user_to from relations where user_id=?) rel on con.Id= rel.user_to", userId,userId).Scan(&contacts)
	return contacts
}
func (p *ChatModel) FindContact(UserId float64) models.Contact {
	contact := models.Contact{}
	p.db.Omit("Relation").Find(&contact, "id = ?", UserId)
	return contact
}
func (p *ChatModel) UpdateMyContact(contact models.Contact, userId float64) models.Contact {
	p.db.Model(&contact).Where("id=?", userId).Updates(models.Contact{Email:contact.Email, Name:contact.Name, Phone:contact.Phone, Password:contact.Password})
	return contact
}

func (p *ChatModel) SaveMessage(message models.Messages) {
	p.db.Create(&message)

}
func (p *ChatModel) SetRelation(relation models.Relations) {
	p.db.Create(&relation)
}
func (p *ChatModel) DeleteRelation(relation models.Relations) {

	p.db.Where("user_id = ? and user_to = ?", relation.UserId, relation.UserTo).Delete(&relation)
}
func (p *ChatModel) ChangeRelation(relation models.Relations) {
	p.db.Model(&relation).Where("user_id = ? and user_to = ?", relation.UserId, relation.UserTo).Update("Relation", relation.Relation)
}
