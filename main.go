package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type User struct {
	login    string
	password string
}

var dataBase []User

func main() {
	router := gin.New()

	router.GET("", MyHandler)
	router.GET("/chop_suey_page", MyHandler1)
	router.GET("/chop_suey_rus_page", MyHandler2)
	router.GET("/toxicity_page", MyHandler3)
	router.GET("/toxicity_rus_page", MyHandler4)

	router.GET("/login_page", PageHandler)
	router.POST("/login", LoginHandler)
	router.POST("/registration", RegistrationHandler)

	router.Run("localhost:8080")
}

func RegistrationHandler(context *gin.Context) {
	user, _ := context.GetQuery("username")
	pas, _ := context.GetQuery("password")

	dataBase = append(dataBase, User{user, pas})

	//fmt.Println(user, pas)
	context.Writer.WriteString("OK")
	//context.Writer.Write([]byte("OK"))
}

func PageHandler(context *gin.Context) {
	html, _ := os.ReadFile("./html/page_with_authorization.html")
	context.Writer.Write(html)
}

func LoginHandler(context *gin.Context) {
	user, _ := context.GetQuery("username")
	pas, _ := context.GetQuery("password")
	fmt.Println(user, pas)
	context.Writer.WriteString("Welcome to the club Body")
	//context.Writer.Write([]byte("OK"))
}

func MyHandler(context *gin.Context) {
	html, _ := os.ReadFile("./html/main_page.html")
	context.Writer.Write(html)
}

func MyHandler1(context *gin.Context) {
	html, _ := os.ReadFile("./html/1page_with_text.html")
	context.Writer.Write(html)
}

func MyHandler2(context *gin.Context) {
	html, _ := os.ReadFile("./html/2page_with_text.html")
	context.Writer.Write(html)
}

func MyHandler3(context *gin.Context) {
	html, _ := os.ReadFile("./html/3page_with_text.html")
	context.Writer.Write(html)
}

func MyHandler4(context *gin.Context) {
	html, _ := os.ReadFile("./html/4page_with_text.html")
	context.Writer.Write(html)
}
