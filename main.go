package main

import (
	"github.com/gin-gonic/gin"
	"main/handler"
)

func main() {
	router := gin.New()

	router.GET("", handler.MyHandler)
	router.GET("/chop_suey_page", handler.MyHandler1)
	router.GET("/chop_suey_rus_page", handler.MyHandler2)
	router.GET("/toxicity_page", handler.MyHandler3)
	router.GET("/toxicity_rus_page", handler.MyHandler4)

	router.GET("/login_page", handler.PageLogHandler)
	router.GET("/login", handler.LoginHandler)
	router.GET("/registration_page", handler.PageRegHandler)
	router.GET("/registration", handler.RegistrationHandler)

	router.GET("/sort_slice_page", handler.SortHandler)
	router.GET("/sort_slice", handler.ArraySortHandler)

	router.Run("localhost:8080")
}
