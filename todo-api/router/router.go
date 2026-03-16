package router

import (
	"database/sql"

	"todo-api/handlers"
	"todo-api/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {

	r := mux.NewRouter()

	authHandler := handlers.AuthHandler{DB: db}
	taskHandler := handlers.TaskHandler{DB: db}

	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	api.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	api.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	api.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	return r
}
