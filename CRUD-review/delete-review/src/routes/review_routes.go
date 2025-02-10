package routes

import (
	"delete-review/src/handlers"

	"github.com/gorilla/mux"
)

// RegisterReviewRoutes registra las rutas del microservicio
func RegisterReviewRoutes(router *mux.Router) {
	router.HandleFunc("/reviews/{id}", handlers.DeleteReview).Methods("DELETE")
}
