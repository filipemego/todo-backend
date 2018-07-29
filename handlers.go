package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func createTodoHandler(service *TodoService) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			todoGetHandler(writer, request, service)
		case "POST":
			todoPostHandler(writer, request, service)
		default:
			writer.WriteHeader(405)
		}
	}
}

func todoGetHandler(writer http.ResponseWriter, request *http.Request, service *TodoService) {
	todoEntries, err := service.Get()
	if err != nil {
		log.Fatal(err)
	}
	jsonWrite(writer, 200, todoEntries)
}

func todoPostHandler(writer http.ResponseWriter, request *http.Request, service *TodoService) {
	todoEntry, err := parseTodoEntry(request)
	if err != nil {
		log.Fatal(err)
	}
	repErr := service.Save(todoEntry)
	if repErr != nil {
		log.Fatal(repErr)
	}

	writer.WriteHeader(201)
}

func parseTodoEntry(request *http.Request) (Todo, error) {
	var entry Todo
	decoder := json.NewDecoder(request.Body)
	return entry, decoder.Decode(&entry)
}

func jsonWrite(writer http.ResponseWriter, status int, content Todos) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	json, jerr := json.Marshal(content)
	if jerr != nil {
		return jerr
	}

	if _, werr := writer.Write(json); werr != nil {
		return werr
	}

	return nil
}
