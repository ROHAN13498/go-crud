package user

import (
	"errors"
)

type UserRepository interface {
	Create(user User) (User, error)
	GetbyId(id int) (User, error)
	List() ([]User, error)
	DeleteUserById(id int) error
}

type InMemoryMap struct {
	userMap map[int]User
}

func NewInMemoryMap() *InMemoryMap {
	return &InMemoryMap{
		userMap: make(map[int]User),
	}
}

func (im *InMemoryMap) Create(user User) (User, error) {
	if im.userMap == nil {
		return User{}, errors.New("User Map is not initialized")
	}

	id := len(im.userMap) + 1
	user.Id = id
	im.userMap[id] = user
	return user, nil
}

func (im *InMemoryMap) GetbyId(id int) (User, error) {
	if id <= 0 {
		return User{}, errors.New("Id is invalid")
	}

	user, ok := im.userMap[id]
	if !ok {
		return User{}, errors.New("User not found")
	}
	return user, nil
}

func (im *InMemoryMap) List() ([]User, error) {
	var users []User = make([]User, 0)
	for _, v := range im.userMap {
		users = append(users, v)
	}
	return users, nil
}

func (im *InMemoryMap) DeleteUserById(id int) error {
	_, ok := im.userMap[id]
	if !ok {
		return errors.New("User Not Found")
	}

	delete(im.userMap, id)
	return nil
}
