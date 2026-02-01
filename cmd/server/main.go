package main

import (
	"log"
	"net/http"
	"github.com/ROHAN13498/go-crud/internal/user"
)

func main(){
	repo:=user.NewInMemoryMap()

	service:= user.NewUserService(repo)

	handler:= user.CreateNewHandler(service)

	http.HandleFunc("/users", func (w http.ResponseWriter,r *http.Request)  {
		switch r.Method {
		case http.MethodPost:
			handler.CreateUser(w,r)
		case http.MethodGet:
			handler.ListAllUsers(w,r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/users/", func(w http.ResponseWriter,r* http.Request){
		switch r.Method{
		case http.MethodGet :
			handler.GetUserById(w,r)
		case http.MethodDelete:
			handler.DeleteUserById(w,r)
		}
	})
	
	log.Println("Server started on :7010")
	log.Fatal(http.ListenAndServe(":7010",nil))
}