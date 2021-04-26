package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"ChatPerson/config"
	"ChatPerson/session"
	"strconv"
	"strings"
	"time"
)

func readHttpBody(response *http.Response) string {
	fmt.Println("Reading body")
	bodyBuffer := make([]byte, 5000)
	var str string
	count, err := response.Body.Read(bodyBuffer)
	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {
		if err != nil {
		}
		str += string(bodyBuffer[:count])
	}
	fmt.Println(str)
	return str
}

func GetAccessToken(client_id string, code string, secret string, callbackUri string) AccessToken {
	fmt.Println("GetAccessToken")
	var token AccessToken
	if callbackUri == "http://localhost:8080/FBLogin" { //For Facebook
		response, err := http.Get("https://graph.facebook.com/oauth/access_token?client_id=" +
			client_id + "&redirect_uri=" + callbackUri +
			"&client_secret=" + secret + "&code=" + code)
		if err == nil {
			auth := readHttpBody(response)
			tokenArr := strings.Split(auth, ",")
			token.Token = strings.Split(tokenArr[0], ":")[1]
			token.Token = token.Token[1 : len(token.Token)-1]
			expire := strings.Split(tokenArr[2], ":")[1]
			expireInt, err := strconv.Atoi(strings.Split(expire, "}")[0])
			if err == nil {
				token.Expiry = int64(expireInt)
			}
		}
	} else { //For Google
		postBody, _ := json.Marshal(map[string]string{
			"code":          code,
			"client_id":     client_id,
			"client_secret": secret,
			"redirect_uri":  callbackUri,
			"grant_type":    "authorization_code",
		})
		responseBody := bytes.NewBuffer(postBody)
		response, err := http.Post("https://accounts.google.com/o/oauth2/token", "application/json", responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer response.Body.Close()
		//	if err == nil {
		auth := readHttpBody(response)
		tokenArr := strings.Split(auth, ",")
		token.Token = strings.Split(tokenArr[0], ":")[1]
		token.Token = token.Token[1 : len(token.Token)-1]
		expire := strings.Split(tokenArr[1], ":")[1]
		expireInt, err := strconv.Atoi(expire)
		if err == nil {
			token.Expiry = int64(expireInt)
		}
	}
	return token
}


func (h *handler) FBLogin(c echo.Context) error {
	code := c.FormValue("code")
	fb := config.GetConfig("fb")
	accessToken := GetAccessToken(fb.ClientID, code, fb.ClientSecret, fb.RedirectURL)
	response, err := http.Get("https://graph.facebook.com/me?access_token=" + accessToken.Token)
	if err != nil {
		fmt.Println(err)
	}
	str := readHttpBody(response)
	user, _ := jason.NewObjectFromBytes([]byte(str))
	name, _ := user.GetString("name")
	sessionId := inMemorySession.Init(name)
	cookie = &http.Cookie{
		Name:    "Token",
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
		MaxAge:  60 * 60,
	}
	c.SetCookie(cookie)
	Customer = h.ChatModel.FindCustomerByName(name)
	if Customer.Name == "" {
		Customer.Name = name
		h.ChatModel.CreateUser(Customer)
		Customer.Id = h.ChatModel.FindCustomerId(name)
	}
	fmt.Println("Endpoint Hit: FBLogin", Customer)
		 return c.Redirect(http.StatusMovedPermanently, "/#/chat")
	//return c.String(http.StatusOK, Customer.Name)
}
func (h *handler) GoogleLogin(c echo.Context) error {
	code := c.FormValue("code")
	google := config.GetConfig("google")
	accessToken := GetAccessToken(google.ClientID, code, google.ClientSecret, google.RedirectURL)
	response, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + accessToken.Token)
	if err != nil {
		fmt.Println(err)
	}
	str := readHttpBody(response)
	user, _ := jason.NewObjectFromBytes([]byte(str))
	email, _ := user.GetString("email")
	sessionId := inMemorySession.Init(email)
	cookie = &http.Cookie{
		Name:    "Token",
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
		MaxAge:  60 * 60,
	}
	c.SetCookie(cookie)
	Customer = h.ChatModel.FindCustomerByName(email)
	if Customer.Name == "" {
		Customer.Email = email
		Customer.Name = email
		h.ChatModel.CreateUser(Customer)
		Customer.Id = h.ChatModel.FindCustomerId(email)
	}
	fmt.Println("Endpoint Hit: GoogleLogin", cookie)
	//return c.String(http.StatusOK, Customer.Email)
	 return c.Redirect(http.StatusMovedPermanently, "/#/chat")
}


func (h *handler) Authorisation(c echo.Context) error {

	inMemorySession = session.NewSession()
	fmt.Println("Endpoint Hit: Authorisation")
	fb := config.GetConfig("fb")
	fbConfig := &oauth2.Config{
		ClientID:     fb.ClientID,
		ClientSecret: fb.ClientSecret,
		RedirectURL:  fb.RedirectURL,
		Scopes:       []string{"email", "user_birthday", "user_location"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}
	FB := fbConfig.AuthCodeURL("")
	google := config.GetConfig("google")
	GoogleConfig := oauth2.Config{
		ClientID:     google.ClientID,
		ClientSecret: google.ClientSecret,
		RedirectURL:  google.RedirectURL,
		Scopes:       []string{"email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	Google := GoogleConfig.AuthCodeURL("")
	return c.JSON(http.StatusOK, map[string]string{
		"FB":     FB,
		"Google": Google,
	})
}


