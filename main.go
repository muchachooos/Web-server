package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.New()

	router.GET("", MyHandler)
	router.GET("/1page_with_text", MyHandler1)
	router.GET("/2page_with_text", MyHandler2)
	router.GET("/3page_with_text", MyHandler3)
	router.GET("/4page_with_text", MyHandler4)

	router.GET("/login_page", pageHandler)
	router.POST("/login", LoginHandler)

	router.Run("localhost:8080")
}
func pageHandler(context *gin.Context) {
	html, _ := os.ReadFile("./html/page_with_authorization.html")
	context.Writer.Write(html)
}

func LoginHandler(context *gin.Context) {
	user, _ := context.GetQuery("username")
	pas, _ := context.GetQuery("password")
	fmt.Println(user, pas)
	context.Writer.WriteString("OK")
	//context.Writer.Write([]byte("OK"))
}

func MyHandler(context *gin.Context) {
	html, _ := os.ReadFile("./html/main_page.html")
	context.Writer.Write(html)
}

func MyHandler1(context *gin.Context) {
	html, _ := os.ReadFile("./html/1page_with_text.html")
	context.Writer.Write(html)
	//t, _ := context.GetQuery("some_key")
	//context.Writer.WriteString("Ebal4" + t + " Ebala")
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
