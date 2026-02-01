package user

import "errors"

type UserService struct{
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService{
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User)(User ,error){
	if user.Name=="" {
		return  User{},errors.New("Invalid Name");
	}
	if user.Email=="" {
		return User{},errors.New("Email is empty");
	}

	return s.repo.Create(user)
}

func (s *UserService) GetUserById(Id int)(User,error){
	if Id<=0{
		return User{},errors.New("Invalid Id")
	}
	return s.repo.GetbyId(Id)
}

func (s *UserService) ListAllUsers()([]User,error)  {
	return s.repo.List();
}

func (s * UserService) DeleteUserById(Id int)error{
	if Id<=0 {
		return errors.New("Invalid Id")
	}
	err:=s.repo.DeleteUserById(Id);

	if err!=nil{
		return err
	}

	return nil;
}