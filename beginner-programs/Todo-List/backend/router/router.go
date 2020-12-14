package router

import (
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todo", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/todo", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/todo/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/todo/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/todo/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/todo", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
