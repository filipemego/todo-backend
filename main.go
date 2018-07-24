package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	todoRepository := NewTodoRepository(Database{})
	http.HandleFunc("/v1/todo", createTodoHandler(todoRepository))

	http.ListenAndServe(":8080", nil)
}

func jsonWrite(writer http.ResponseWriter, content Todos) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	json, jerr := json.Marshal(content)
	if jerr != nil {
		return jerr
	}

	if _, werr := writer.Write(json); werr != nil {
		return werr
	}

	return nil
}
