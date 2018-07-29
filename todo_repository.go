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

// Get returns all Todo from DB.
func (repository *TodoRepository) Get(id string) (Todos, error) {
	var fetchedID, title, content, created, modified string
	err := repository.db.Select("SELECT * FROM todos LIMIT 1", &fetchedID, &title, &content, &created, &modified)
	if err != nil {
		return Todos{}, err
	}
	todos := Todos{
		Todo{
			ID:       fetchedID,
			Title:    title,
			Content:  content,
			Created:  created,
			Modified: modified,
		},
	}
	return todos, nil
}

// Save store a record on the database.
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
