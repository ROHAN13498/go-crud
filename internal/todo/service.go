package todo

import (
	"errors"
)

type UserGetter interface {
	UserExists(id int) error
}

type TodoService struct {
	repo        TodoRepo
	UserService UserGetter
}

func NewTodoSevice(repo TodoRepo, userService UserGetter) *TodoService {
	return &TodoService{repo: repo, UserService: userService}
}

func (s *TodoService) CreateTodo(userId int, title string) (Todo, error) {
	if title == "" {
		return Todo{}, errors.New("Empty User Id is provided")
	}

	if err := s.UserService.UserExists(userId); err != nil {
		return Todo{}, errors.New("Invalid User")
	}

	return s.repo.Create(Todo{
		UserId: userId,
		Title:  title,
		Done:   false,
	})
}

func (s *TodoService) ListTodos(userId int) ([]Todo, error) {
	return s.repo.ListByUserId(userId)
}

func (s *TodoService) DeleteTodo(taskId int) error {
	return s.repo.DelteTask(taskId)
}
