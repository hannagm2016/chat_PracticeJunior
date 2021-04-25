package handlers

import (
	"ChatPerson/models"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

var Customers []models.Contact


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
