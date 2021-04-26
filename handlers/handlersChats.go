package handlers

import (
	"ChatPerson/models"
	"ChatPerson/repository"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

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