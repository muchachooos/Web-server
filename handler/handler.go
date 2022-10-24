package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/sort"
	"os"
	"strconv"
)

type User struct {
	login    string
	password string
}

var dataBase []User

func ArraySortHandler(context *gin.Context) {
	values := context.Request.URL.Query()

	// достаём срез из строк
	slice := values["array"]

	fmt.Println("11111", slice)

	var arr []int

	var val int

	for i := range slice {
		val, _ = strconv.Atoi(slice[i]) // превращаем строку в число
		arr = append(arr, val)
	}

	sorted := sort.BubbleSort(arr)

	fmt.Println("AAAAAAAA", sorted)

	context.Writer.WriteString("Bubble sort: " + fmt.Sprint(sorted))

	fmt.Println("slice = ", slice)

	fmt.Println(sort.BubbleSort(arr))
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

func SortHandler(context *gin.Context) {
	html, _ := os.ReadFile("./html/sort_slice_page.html")
	context.Writer.Write(html)
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
