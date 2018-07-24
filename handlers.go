package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func createTodoHandler(repository *TodoRepository) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		queryString, _ := url.ParseQuery(request.URL.RawQuery)
		fmt.Printf("%v\n", queryString)
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
