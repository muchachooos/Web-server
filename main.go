package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"main/handler"
)

func main() {
	router := gin.New()

	//Подключаемся к SQL и DB
	dataBase, err := sqlx.Open("mysql", "root:040498usa_wot@tcp(127.0.0.1:3306)/userdata")
	if err != nil {
		panic(err)
		return
	}

	server := handler.Server{
		dataBase,
	}

	router.GET("", handler.MyHandler)
	router.GET("/chop_suey_page", handler.MyHandler1)
	router.GET("/chop_suey_rus_page", handler.MyHandler2)
	router.GET("/toxicity_page", handler.MyHandler3)
	router.GET("/toxicity_rus_page", handler.MyHandler4)

	router.GET("/login_page", handler.PageLogHandler)
	router.GET("/login", server.LoginHandler)
	router.GET("/registration_page", handler.PageRegHandler)
	router.GET("/registration", server.RegistrationHandler)

	router.GET("/sort_slice_page", handler.SortHandler)
	router.GET("/sort_slice", handler.ArraySortHandler)

	router.Run("localhost:8080")
}

/*
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
			context.Writer.WriteString("This login or password already exist. Try again")
			return
		}
		if pass == dataBase[i].password {
			context.Writer.WriteString("This password already exist. Try again")
			return
		}
	}

	dataBase = append(dataBase, User{user, pass})

	context.Writer.WriteString("Welcome to the club Body")
}
*/

//a, ok := context.GetQuery("username") //Достаём Query-параметр(user = key(username))
//if a == "" || !ok {                   //ok == false; Поверка на пустые значения
//	context.Writer.WriteString("No username")
//	return
//}
//
//b, ok := context.GetQuery("password") //Достаём Query-параметр(pass = key(password))
//if b == "" || !ok {                   //ok == false; Поверка на пустые значения
//	context.Writer.WriteString("No password")
//	return
//}

//log
//
//user, ok := context.GetQuery("username")
//if user == "" || !ok { //ok == false; Поверка на пустые значения
//	context.Writer.WriteString("No username")
//	return
//}
//
//pass, ok := context.GetQuery("password")
//if pass == "" || !ok { //ok == false; Поверка на пустые значения
//	context.Writer.WriteString("No password")
//	return
//}
//
//for i := range dataBaseArr {
//	if user == dataBaseArr[i].login && pass == dataBaseArr[i].password {
//		fmt.Println(user, pass)
//		context.Writer.WriteString("Welcome to the club Body")
//		return
//	}
//}

//type User struct {
//	login    string
//	password string
//}
//
//var dataBaseArr []User
