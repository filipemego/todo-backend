package main

// TodoService struct.
type TodoService struct {
	repository *TodoRepository
}

// NewTodoService TodoService constructor.
func NewTodoService(repository *TodoRepository) *TodoService {
	return &TodoService{repository: repository}
}

// Get returns all records from the DB.
func (service *TodoService) Get() (Todos, error) {
	todos, err := service.repository.Get("260954e3-8a3c-41d9-87c4-371db1241e83")
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Save store a record on the database.
func (service *TodoService) Save(entry Todo) error {
	_, err := service.repository.Save(entry)
	return err
}
