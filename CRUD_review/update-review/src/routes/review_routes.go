package routes

import (
	"update-review/src/handlers"

	"github.com/gorilla/mux"
)

// RegisterReviewRoutes configura las rutas del microservicio
func RegisterReviewRoutes(router *mux.Router) {
	router.HandleFunc("/api/reviews/{id}", handlers.UpdateReview).Methods("PUT")
}
