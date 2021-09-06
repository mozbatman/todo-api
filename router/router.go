package router

import (
	"server/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	return router
}
