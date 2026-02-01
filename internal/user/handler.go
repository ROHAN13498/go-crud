package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct{
	service *UserService
}

func CreateNewHandler(service *UserService) *Handler{
	return &Handler{service:service}
}

func (h* Handler) CreateUser(w http.ResponseWriter, r *http.Request){
	var req CreateUserRequest

	decoder:=json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err:= decoder.Decode(&req);
	if err!=nil {
		log.Print(err)
		http.Error(w,"Invalid Request Body",http.StatusBadRequest);
		return;
	}
	user:= User{
		Name: req.Name,
		Email: req.Email,
	}
	newUser,err:= h.service.CreateUser(user)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	resp:= CreateUserResponse{
		Id: newUser.Id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h* Handler) GetUserById(w http.ResponseWriter, r *http.Request){
	idStr:=strings.TrimPrefix(r.URL.Path,"/users/");
	id,err:=strconv.Atoi(idStr)

	if(err!=nil){
		http.Error(w,"Invalid Number Value",http.StatusBadRequest)
		return ;
	}

	user,err:=h.service.GetUserById(id);

	if(err!=nil){
		http.Error(w,"Invalid Number Value",http.StatusBadRequest)
		return ;
	}

	resp:=CreateUserRequest{
		Id:user.Id,
		Name:user.Name,
		Email: user.Email,
	}


	w.Header().Set("Content-Type","application/json");
	json.NewEncoder(w).Encode(resp)
}

func (h * Handler) ListAllUsers(w http.ResponseWriter,r* http.Request){
	users,err:=h.service.ListAllUsers()
	if(err!=nil){
		http.Error(w,"Something went wrong",http.StatusBadRequest)
		return
	}

	resp:=make([]UserResponse,0,len(users))

	for _,u:= range users{
		resp=append(resp, UserResponse{
			Id: u.Id,
			Name: u.Name,
		})
	}

	w.Header().Set("Content-Type","application/json");
	json.NewEncoder(w).Encode(resp)
}

func (h* Handler) DeleteUserById(w http.ResponseWriter,r *http.Request){
	idStr:=strings.TrimPrefix(r.URL.Path,"/users/");
	id,err:=strconv.Atoi(idStr)
	fmt.Print(id)
	if(err!=nil){
		http.Error(w,err.Error(),http.StatusBadRequest)
		return ;
	}

	deleteError:=h.service.DeleteUserById(id)


	if deleteError!=nil {
		http.Error(w,deleteError.Error(),http.StatusBadRequest)
		return;
	}
	w.WriteHeader(http.StatusNoContent)
}


