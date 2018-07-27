package main

import (
	"log"
	"net/http"
)

func main() {
	database := NewDatabase("todo_app:1234pw@tcp(0.0.0.0:3306)/todos")
	todoRepository := NewTodoRepository(database)
	todoService := NewTodoService(todoRepository)
	defer database.Close()

	http.HandleFunc("/v1/todo", createTodoHandler(todoService))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
