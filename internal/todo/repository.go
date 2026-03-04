package todo

import "errors"

type TodoRepo interface {
	Create(todo Todo) (Todo, error)
	ListByUserId(userId int) ([]Todo, error)
	DelteTask(id int) error
}

type InMemoryTodoMap struct {
	data map[int]Todo
}

func NewTodoMap() *InMemoryTodoMap {
	return &InMemoryTodoMap{
		data: make(map[int]Todo),
	}
}

func (r *InMemoryTodoMap) Create(t Todo) (Todo, error) {
	id := len(r.data) + 1
	t.Id = id
	r.data[id] = t
	return t, nil
}

func (r *InMemoryTodoMap) ListByUserId(userId int) ([]Todo, error) {
	res := []Todo{}
	for _, t := range r.data {
		if t.UserId == userId {
			res = append(res, t)
		}
	}
	return res, nil
}

func (r *InMemoryTodoMap) DelteTask(id int) error {
	if _, ok := r.data[id]; !ok {
		return errors.New("TaskId not found")
	}
	delete(r.data, id)
	return nil
}
