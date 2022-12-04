package handler

import (
	"Web-server/sort"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SortHandler(context *gin.Context) {

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
