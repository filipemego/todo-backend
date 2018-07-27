package main

// TodoService struct
type TodoService struct {
	repository *TodoRepository
}

// NewTodoService TodoService constructor
func NewTodoService(repository *TodoRepository) *TodoService {
	return &TodoService{repository: repository}
}

// Save store a record on the database
func (repository *TodoService) Save(entry Todo) {
	lastId, err := repository.Save(entry)
}
