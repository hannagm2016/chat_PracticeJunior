package main

import (
	"ChatPerson/db"
	"github.com/labstack/echo"
	//	"site/models"
	"ChatPerson/handlers"
	"ChatPerson/repository"
	"github.com/labstack/echo/middleware"
	//"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	d := db.DBConnect()
	h := handlers.NewHandler(repository.NewChatModel(d))
	e.Static("/", "./chat_vue/dist/")

	e.GET("/chats", h.Chats)
	e.GET("/chat/:userId", h.Chat)
	e.GET("/contacts", h.Contacts)
	//e.GET("/chat/:name", h.SingleChat)
	e.POST("/message", h.AddMessage)
	e.GET("/contact/:id", h.SingleCustomer)
	e.GET("/socket", socketReaderCreate)

	e.Logger.Fatal(e.Start(":8080"))
}

var upgrader = websocket.Upgrader{}

//socketReader struct
type socketReader struct {
	con *websocket.Conn
}
type Mes struct {
	// Name string `json:"name"`
	Text string //`json:"text"`
	Time string
	Type string
	//  UserId float64
}

var savedsocketreader []*socketReader

func socketReaderCreate(c echo.Context) error {
	log.Println("socket request")
	con, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	ptrSocketReader := &socketReader{
		con: con,
	}
	savedsocketreader = append(savedsocketreader, ptrSocketReader)
	ptrSocketReader.startThread()
	return err
}

func (i *socketReader) broadcast(str string) {
	for _, g := range savedsocketreader {
		if g == i { // no send message to himself
			continue
		}
		fmt.Println(g, "broadcast", i)
		g.writeMsg(str)
	}
}

func (i *socketReader) read() {
	_, b, er := i.con.ReadMessage()
	if er != nil {
		panic(er)
	}
	log.Println(string(b), "read")
	i.broadcast(string(b))
}

func (i *socketReader) writeMsg(str string) {
	message := Mes{Text: str}
	msg, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(msg), "writeMsg")
	i.con.WriteJSON(string(msg))
}

func (i *socketReader) startThread() {

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(err, "***")
			}
			log.Println("thread socketreader finish")
		}()

		for {
			i.read()
		}
	}()
}
