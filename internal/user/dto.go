package user


type CreateUserRequest struct{
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}


type CreateUserResponse struct {
	Id int
}

type UserResponse struct{
	Id int `json:"id"`
	Name string `json:"name"`
}
