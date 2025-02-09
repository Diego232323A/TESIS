package routes

import (
	"create-review/src/handlers"

	"github.com/gorilla/mux"
)

// Registra las rutas del microservicio
func RegisterReviewRoutes(router *mux.Router) {
	router.HandleFunc("/review", handlers.CreateReviewHandler).Methods("POST")
}
