package handlers

import (

	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"ChatPerson/models"
	"ChatPerson/session"
	"strconv"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)
const SecretKey ="secret"
var inMemorySession *session.Session

var Customer models.Contact

var cookie *http.Cookie

type AccessToken struct {
	Token  string
	Expiry int64
}

func (h *handler) Login(c echo.Context) error {
    var data map[string]string
        if err:=c.Bind (&data); err !=nil {
            return err
        }
        fmt.Println(data)
        var user models.Contact
    user = h.ChatModel.FindCustomerByName(data["Name"])
     if user.Name == "" {
     fmt.Println("User not found")
        return c.JSON(http.StatusNotFound, "User not found")
    }
  /*  if err :=bcrypt.CompareHashAndPassword(user.Password,[]byte(data["Password"])); err!=nil {
            fmt.Println("Incorrect Password")
        return c.JSON(http.StatusBadRequest, "Incorrect Password")
    }*/
Customer=user;
    claims :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
        Issuer: strconv.Itoa(int(user.Id)),
        ExpiresAt: time.Now().Add(time.Hour*24).Unix(),
    })
    token, err :=claims.SignedString([]byte(SecretKey))
    if err !=nil {
        return c.String(http.StatusInternalServerError, "Could not login")
    }
        cookie:=http.Cookie{
            Name: "jwt",
            Value: token,
            Expires: time.Now().Add(time.Hour *24),
            HttpOnly: true,
        }
        c.SetCookie(&cookie)
        fmt.Println(cookie,"*")
    return c.JSON (http.StatusOK, token)
}

func (h *handler) User(c echo.Context) error {
    cookie, err :=c.Cookie("jwt")
    if err!=nil {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface {}, error) {
    return []byte(SecretKey), nil
    })
    if err!=nil {
        return c.JSON(http.StatusUnauthorized, "Unauthorized")
    }
    claims := token.Claims.(*jwt.StandardClaims)
    id, _ :=strconv.Atoi(claims.Issuer)
    user := h.ChatModel.FindCustomerById(id)
    return c.JSON(http.StatusOK, user)
}

func (h *handler) Logout(c echo.Context) error {
    cookie:= http.Cookie {
        Name: "jwt",
        Value: "",
        Expires: time.Now().Add(-time.Hour),
        HttpOnly: true,
    }
    c.SetCookie(&cookie)
       var no  models.Contact
        Customer= no
        fmt.Println(Customer,"+++", cookie)
         return c.Redirect(http.StatusMovedPermanently, "/")
        //return c.String (http.StatusOK, "Logged Outed")
}

func (h *handler) RegistrationPost(c echo.Context) error {
	var data map[string]string
if err:=c.Bind (&data); err !=nil {
    return err
}
	fmt.Println("Endpoint Hit: SaveUser", data)
	password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := models.Contact {
        Name: data["Name"],
        Email: data["Email"],
        Password: password,
	}
	h.ChatModel.CreateUser(user)
	return c.JSON(http.StatusOK, user)
}

