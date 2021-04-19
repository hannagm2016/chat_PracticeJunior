package handlers

import (
	"ChatPerson/models"
	"ChatPerson/repository"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

var Customers []models.Contact
var Chats []models.Chats

type handler struct {
	ChatModel repository.ChatModelImpl
}

func NewHandler(p repository.ChatModelImpl) *handler {
	return &handler{p}
}

func (h *handler) Chats(c echo.Context) error {
	Chats = h.ChatModel.FindChats(Customer.Id)
	fmt.Println("Endpoint Hit: All chats")
	return c.JSON(http.StatusOK, Chats)
}

func (h *handler) Chat(c echo.Context) error {
	id := c.Param("userId")
	key, _ := strconv.ParseFloat(string(id), 64)
	chat := h.ChatModel.FindChat(key)
	fmt.Println("Endpoint Hit: ReturnChat with user - ", id)
	return c.JSON(http.StatusOK, chat)
}
func (h *handler) Contacts(c echo.Context) error {
	Customers = h.ChatModel.FindContacts()
	fmt.Println("Endpoint Hit: All customers")

	return c.JSON(http.StatusOK, Customers)
}

func (h *handler) SingleCustomer(c echo.Context) error {
	id := c.Param("id")
	key, _ := strconv.ParseFloat(string(id), 64)
	contact := h.ChatModel.FindContact(key)
	fmt.Println("Endpoint Hit: ReturnCustomer details- ", id)
	return c.JSON(http.StatusOK, contact)
}

func (h *handler) AddMessage(c echo.Context) error {
	var message models.Message
	if err := c.Bind(&message); err != nil {
		return err
	}
	mes := models.Messages{UserFromId: Customer.Id, UserToId: message.UserId, Time: message.Time, Text: message.Text}

	fmt.Println("Endpoint Hit: SaveMessage", mes)
	h.ChatModel.SaveMessage(mes)

	return c.JSON(http.StatusOK, message)
}
