package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/sort"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DataBase *sqlx.DB
}

type Data struct {
	ID       int    `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

func (s *Server) RegistrationHandler(context *gin.Context) {

	var err error

	a, ok := context.GetQuery("username") //Достаём Query-параметр(a = key(username))
	if a == "" || !ok {                   //ok == false; Поверка на пустые значения
		context.Writer.WriteString("No username")
		return
	}

	b, ok := context.GetQuery("password") //Достаём Query-параметр(b = key(password))
	if b == "" || !ok {                   //ok == false; Поверка на пустые значения
		context.Writer.WriteString("No password")
		return
	}

	_, err = s.DataBase.Exec("INSERT INTO users(login, password) VALUES (?,?)", a, b) //Добавляем значения в БД
	if err != nil {
		context.Writer.WriteString("This login already exist. Try again")
		context.Status(500)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) LoginHandler(context *gin.Context) {

	var err error

	log, ok := context.GetQuery("username")
	if log == "" || !ok { //ok == false; Поверка на пустые значения
		context.Writer.WriteString("No username")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok { //ok == false; Поверка на пустые значения
		context.Writer.WriteString("No password")
		return
	}

	var resultTable []Data

	//Возвращаем значение по логину(с) и паролю(d) или ошибку
	err = s.DataBase.Select(&resultTable, "SELECT * FROM users WHERE login = ? AND password = ?", log, pass)
	if err != nil {
		context.Writer.WriteString("Wrong login or password. Try again")
		context.Status(500)
		return
	}

	if len(resultTable) == 0 {
		context.Writer.WriteString("Wrong login or password. Try again")
		return
	}

	fmt.Println(resultTable)

	context.Writer.WriteString("Welcome to the club Body")
}

func ArraySortHandler(context *gin.Context) {
	values := context.Request.URL.Query()

	// достаём срез из строк
	slice := values["array"]

	var arr []int
	var val int
	var err error

	for i := range slice {
		val, err = strconv.Atoi(slice[i])
		if err != nil {
			context.Writer.WriteString("Error! Contains string")
			return
		}
	}

	for i := range slice {
		val, _ = strconv.Atoi(slice[i]) // превращаем строку в число
		arr = append(arr, val)
	}

	sorted := sort.BubbleSort(arr)

	context.Writer.WriteString("Bubble sort: " + fmt.Sprint(sorted))
}

func SortHandler(context *gin.Context) {
	html, _ := os.ReadFile("./resources/html/sort_slice_page.html")
	context.Writer.Write(html)
}

func PageRegHandler(context *gin.Context) {
	html, _ := os.ReadFile("./resources/html/page_with_registration.html")
	context.Writer.Write(html)
}

func PageLogHandler(context *gin.Context) {
	html, _ := os.ReadFile("./resources/html/page_with_authorization.html")
	context.Writer.Write(html)
}

func MyHandler(context *gin.Context) {
	html, _ := os.ReadFile("./resources/html/main_page.html")
	context.Writer.Write(html)
}

func MyHandler1(context *gin.Context) {
	html, _ := os.ReadFile("./resources/html/1page_with_text.html")
	context.Writer.Write(html)
}

func MyHandler2(context *gin.Context) {
	html, _ := os.ReadFile("./resources/html/2page_with_text.html")
	context.Writer.Write(html)
}

func MyHandler3(context *gin.Context) {
	html, _ := os.ReadFile("./resources/html/3page_with_text.html")
	context.Writer.Write(html)
}

func MyHandler4(context *gin.Context) {
	html, _ := os.ReadFile("./resources/html/4page_with_text.html")
	context.Writer.Write(html)
}
