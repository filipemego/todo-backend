package main

// Todos list of Todo
type Todos []Todo

// Todo struct
type Todo struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
