package routes

import (
	"geo-service/src/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/places", handlers.CreatePlace).Methods("POST")
	router.HandleFunc("/places/nearby", handlers.GetNearbyPlaces).Methods("GET")
}
