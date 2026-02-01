package user

import (
	"errors"
)
type UserRepository interface {
	Create(user User) (User, error)
	GetbyId(id int)(User,error)
	List()([]User,error)
	DeleteUserById(id int) (error)
}

type InMemoryMap struct{
	userMap map[int]User
}

func NewInMemoryMap () *InMemoryMap{
	return &InMemoryMap{
		userMap: make(map[int]User),
	}
}

func (im *InMemoryMap) Create(user User)(User,error){
	if(im.userMap==nil){
		return User{},errors.New("User Map is not initialized")
	}

	im.userMap[len(im.userMap)+1]=User{Id:len(im.userMap)+1,Name:"Test",Email: "test@gmail.com"}
	return im.userMap[len(im.userMap)], nil
}

func(im *InMemoryMap) GetbyId(id int) (User,error){
	if id<=0 {
		return User{},errors.New("Id is invalid");
	}

	return im.userMap[id],nil
}

func (im *InMemoryMap) List()([]User,error){
	var users []User= make([]User, 0)
	for _,v:=range(im.userMap){
		users=append(users, v)
	}
	return users,nil
}

func (im* InMemoryMap) DeleteUserById(id int) error {
	_,ok:=im.userMap[id];
	if !ok {
		return errors.New("User Not Found")
	}

	delete(im.userMap,id)
	return nil;
}