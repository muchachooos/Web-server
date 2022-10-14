package main

import (
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

	router.Run("localhost:8080")
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
