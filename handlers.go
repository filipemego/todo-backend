package main

import (
	"encoding/json"
	"net/http"
)

func createTodoHandler(repository *TodoRepository) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// queryString, _ := url.ParseQuery(request.URL.RawQuery)

		switch request.Method {
		case "GET":
			todoGetHandler(writer, request, repository)
		case "POST":
			todoPostHandler(writer, request, repository)
		}
	}
}

func todoGetHandler(writer http.ResponseWriter, request *http.Request, repository *TodoRepository) {
	todos := repository.GetAll()
	jsonWrite(writer, todos)
}

func todoPostHandler(writer http.ResponseWriter, request *http.Request, repository *TodoRepository) {
	writer.WriteHeader(201)
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
