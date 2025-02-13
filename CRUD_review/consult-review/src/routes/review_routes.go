package routes

import (
	"consult-review/src/handlers"

	"github.com/gorilla/mux"
)

// RegisterReviewRoutes configura las rutas del microservicio
func RegisterReviewRoutes(router *mux.Router) {
	router.HandleFunc("/reviews", handlers.GetAllReviews).Methods("GET")
	router.HandleFunc("/reviews/{id}", handlers.GetReviewByID).Methods("GET")
	router.HandleFunc("/reviews/place/{place_id}", handlers.GetReviewsByPlace).Methods("GET")
}
