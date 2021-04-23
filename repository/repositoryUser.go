package repository
import (
    "ChatPerson/models"
)

func (p *ChatModel) CreateUser(contact models.Contact) {
	p.db.Create(&contact)
}

func (p *ChatModel) FindCustomerByName(userName string) models.Contact {
	var contact models.Contact
	p.db.Find(&contact, "name = ?", userName)
	return contact
}
func (p *ChatModel) FindCustomerByEmail(email string) models.Contact {
	var contact models.Contact
	p.db.Find(&contact, "email = ?", email)
	return contact
}
func (p *ChatModel) FindCustomerById(id int) models.Contact {
	var contact models.Contact
	p.db.Find(&contact, "id = ?", id)
	return contact
}

func (p *ChatModel) FindCustomerId(userName string) float64 {
	var id float64
	p.db.Table("users").Select("id").Find(&id, "name = ?", userName)
	return id
}
