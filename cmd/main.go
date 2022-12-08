package main

import (
	"Web-server/handler"
	"Web-server/model"
	"Web-server/storage"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
)

func main() {
	router := gin.Default()

	var conf model.Config

	fileInBytes, err := os.ReadFile("../../configuration.json")
	if err != nil {
		fmt.Println("Error Read File:", err)
		return
	}

	err = json.Unmarshal(fileInBytes, &conf)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}

	dataBase, err := sqlx.Open("mysql", conf.DataSourceName)
	if err != nil {
		panic(err)
		return
	}

	if dataBase == nil {
		fmt.Println("dB nil")
		panic(err)
		return
	}

	server := handler.Server{
		Storage: &storage.UserStorage{
			DataBase: dataBase,
		},
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

	port := ":" + strconv.Itoa(conf.Port)

	err = router.Run(port)
	if err != nil {
		panic(err)
		return
	}
}
