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
id := c.Param("userId")
key, _ := strconv.ParseFloat(string(id), 64)
	Chats = h.ChatModel.FindChats(key)
	fmt.Println("Endpoint Hit: All chats for", key)
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
    id := c.Param("userId")
    	key, _ := strconv.ParseFloat(string(id), 64)
	Customers = h.ChatModel.FindContacts(key)
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
id := c.Param("userId")
	key, _ := strconv.ParseFloat(string(id), 64)
	var message models.Message
	if err := c.Bind(&message); err != nil {
		return err
	}
	mes := models.Messages{UserFromId: key, UserToId: message.UserId, Time: message.Time, Text: message.Text}

	fmt.Println("Endpoint Hit: SaveMessage", mes, key)
	h.ChatModel.SaveMessage(mes)

	return c.JSON(http.StatusOK, message)
}
func (h *handler) SetRelation(c echo.Context) error {
	var rel models.Relations
	if err := c.Bind(&rel); err != nil {
    		fmt.Println(err)
    		return err
    	}
    		fmt.Println("Endpoint Hit: SetRelation", rel)


	h.ChatModel.SetRelation(rel)

	return c.JSON(http.StatusOK, rel)
}
func (h *handler) ChangeRelation(c echo.Context) error {
	var rel models.Relations
	if err := c.Bind(&rel); err != nil {
    		return err
    	}
    		fmt.Println("Endpoint Hit: ChangeRelation", rel)
	h.ChatModel.ChangeRelation(rel)
	return c.JSON(http.StatusOK, rel)
}
func (h *handler) DeleteRelation(c echo.Context) error {
	var rel models.Relations
	if err := c.Bind(&rel); err != nil {
    		return err
    	}
    		fmt.Println("Endpoint Hit: DeleteRelation", rel)
	h.ChatModel.DeleteRelation(rel)
	return c.JSON(http.StatusOK, rel)
}
