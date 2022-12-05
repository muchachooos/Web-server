package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func PageLogHandler(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/page_with_authorization.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func PageRegHandler(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/page_with_registration.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func PageDelHandler(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/delete_user_page.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func PageChangeHandler(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/change_pass_page.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func PageSortHandler(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/sort_slice_page.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func MyHandler(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/main_page.html")
	if err != nil {
		context.Writer.WriteString("AAAAAAAAAAAA")
		fmt.Println("AAAAAAAAAA:", err)
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func MyHandler1(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/1page_with_text.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func MyHandler2(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/2page_with_text.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func MyHandler3(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/3page_with_text.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}

func MyHandler4(context *gin.Context) {
	html, err := os.ReadFile("./resources/html/4page_with_text.html")
	if err != nil {
		context.Status(500)
		return
	}

	context.Writer.Write(html)
}
