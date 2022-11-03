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

	router.GET("/delete_user_page", handler.PageDelHandler)
	router.GET("/delete_user", server.DeleteHandler)

	router.GET("/change_pass_page", handler.PageChangeHandler)
	router.GET("/change_pass", server.ChangeHandler)

	router.GET("/sort_slice_page", handler.PageSortHandler)
	router.GET("/sort_slice", handler.SortHandler)

	router.Run("localhost:8080")
}
