package main

import (
	"fmt"
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

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	fmt.Println("New todo added: ", newTodo)
	fmt.Println("All todos: ", todos)
	context.IndentedJSON(http.StatusCreated, todos)

}
func getTodoById(id string) (*todo, error) {
	for _, t := range todos {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("todo with ID %s not found", id)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	t, err := getTodoById(id)
	fmt.Println("Todo: ", t)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, t)
}
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.Run("localhost:8080")

}
