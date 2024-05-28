package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Title: "Learn Go", Completed: false},
	{ID: "2", Title: "Build a RESTful API in Go", Completed: false},
	{ID: "3", Title: "Build a React app", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8080")

}
