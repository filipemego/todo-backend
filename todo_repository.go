package main

import (
	"log"

	"github.com/satori/go.uuid"
)

// TodoRepository struct.
type TodoRepository struct {
	db *Database
}

// NewTodoRepository constructor.
func NewTodoRepository(db *Database) *TodoRepository {
	return &TodoRepository{db: db}
}

// GetAll returns all Todo from DB.
func (repository *TodoRepository) GetAll() Todos {
	todos := []Todo{
		Todo{ID: "uuid-1", Title: "Todo 1", Content: "Content 1"},
		Todo{ID: "uuid-2", Title: "Todo 2", Content: "Content 2"},
	}
	return todos
}

// Save store a record on the database
func (repository *TodoRepository) Save(todo Todo) (int64, error) {
	uuid, uErr := uuid.NewV4()
	if uErr != nil {
		log.Fatal(uErr)
	}
	return repository.db.Insert(
		"INSERT INTO todos (id, title, content) VALUE(?, ?, ?)",
		uuid,
		todo.Title,
		todo.Content,
	)
}
