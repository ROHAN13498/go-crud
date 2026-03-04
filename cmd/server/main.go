package main

import (
	"log"
	"net/http"

	"github.com/ROHAN13498/go-crud/internal/httputil"
	"github.com/ROHAN13498/go-crud/internal/todo"
	"github.com/ROHAN13498/go-crud/internal/user"
)

func main() {
	// USER
	repo := user.NewInMemoryMap()
	service := user.NewUserService(repo)
	handler := user.CreateNewHandler(service)
	todoRepo := todo.NewTodoMap()
	todoService := todo.NewTodoSevice(todoRepo, service)
	todoHandler := todo.NewTodoHandler(todoService)

	h := &httputil.Handler{}

	// User routes
	h.Post("/users", handler.CreateUser)
	h.Get("/users", handler.ListAllUsers)
	h.Get("/users/{id}", handler.GetUserById)
	h.Delete("/users/{id}", handler.DeleteUserById)
	// Todo routes
	h.Post("/users/{id}/todos", todoHandler.CreateTodo)
	h.Get("/todos/{id}", todoHandler.GetTodoById)
	h.Delete("/todos/{id}", todoHandler.DeleteTodoById)

	// Debug: catch-all to see what requests are coming in
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("UNMATCHED: method=%q path=%q", r.Method, r.URL.Path)
		http.Error(w, "not found", http.StatusNotFound)
	})

	log.Println("Server started on :7010")
	log.Fatal(http.ListenAndServe(":7010", nil))
}
