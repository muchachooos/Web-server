package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	router := gin.New()

	router.GET("", MyHandler)
	router.GET("/chop_suey_page", MyHandler1)
	router.GET("/chop_suey_rus_page", MyHandler2)
	router.GET("/toxicity_page", MyHandler3)
	router.GET("/toxicity_rus_page", MyHandler4)

	router.GET("/login_page", PageLogHandler)
	router.GET("/login", LoginHandler)
	router.GET("/registration_page", PageRegHandler)
	router.GET("/registration", RegistrationHandler)

	router.GET("/sort", ArraySortHandler)

	router.Run("localhost:8080")
}

type User struct {
	login    string
	password string
}

var dataBase []User

func ArraySortHandler(context *gin.Context) {
	values := context.Request.URL.Query()

	// достаём срез из строк
	slice := values["array"]

	// превращаем строку в число
	integer, _ := strconv.Atoi(slice[0])
	integer1, _ := strconv.Atoi(slice[1])
	integer2, _ := strconv.Atoi(slice[2])

	arr := []int{integer, integer1, integer2}

	context.Writer.WriteString("Welcome to the club Body")

	fmt.Println("int = ", integer)

	fmt.Println("slice = ", slice)

	fmt.Println("val = ", values)

	fmt.Println(bubbleSort(arr))
}

func bubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				fmt.Println(slice)
				slice[j], slice[j+1] = slice[j+1], slice[j]
				fmt.Println(slice)
			}
		}
	}
	return slice
}

func RegistrationHandler(context *gin.Context) {

	user, ok := context.GetQuery("username") //Достаём Query-параметр(user = key(username))
	if user == "" || !ok {                   //ok == false; Поверка на пустые значения
		context.Writer.WriteString("No username")
		return
	}

	pass, ok := context.GetQuery("password") //Достаём Query-параметр(pass = key(password))
	if pass == "" || !ok {                   //ok == false; Поверка на пустые значения
		context.Writer.WriteString("No password")
		return
	}

	for i := range dataBase {
		if user == dataBase[i].login {
			context.Writer.WriteString("This login already exist. Try again")
			return
		}
		if pass == dataBase[i].password {
			context.Writer.WriteString("This password already exist. Try again")
			return
		}
	}

	dataBase = append(dataBase, User{user, pass})

	context.Writer.WriteString("Welcome to the club Body")
	//context.Writer.Write([]byte("OK"))

	fmt.Println("-------RegistrationHandler---------")
	fmt.Println(dataBase)
}
func LoginHandler(context *gin.Context) {

	user, ok := context.GetQuery("username")
	if user == "" || !ok { //ok == false; Поверка на пустые значения
		context.Writer.WriteString("No username")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok { //ok == false; Поверка на пустые значения
		context.Writer.WriteString("No password")
		return
	}

	for i := range dataBase {
		if user == dataBase[i].login && pass == dataBase[i].password {
			fmt.Println(user, pass)
			context.Writer.WriteString("Welcome to the club Body")
			return
		}
	}

	fmt.Println("-------LoginHandler---------")
	fmt.Println(user, pass)
	context.Writer.WriteString("Wrong login or password. Try again")
	//context.Writer.Write([]byte("OK"))
}

func PageRegHandler(context *gin.Context) {
	html, _ := os.ReadFile("./html/page_with_registration.html")
	context.Writer.Write(html)
}

func PageLogHandler(context *gin.Context) {
	html, _ := os.ReadFile("./html/page_with_authorization.html")
	context.Writer.Write(html)
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
