package main

import (
	"net/http"
)

func main() {
	database := NewDatabase("mysql/db")
	todoRepository := NewTodoRepository(database)
	defer database.Close()

	http.HandleFunc("/v1/todo", createTodoHandler(todoRepository))

	http.ListenAndServe(":8080", nil)
}
