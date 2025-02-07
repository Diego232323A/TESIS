package main

import (
	"create-review/src/config"
	_ "create-review/src/docs"
	"create-review/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Microservicio de Creación de Reseñas
// @version 1.0
// @description API para agregar reseñas de lugares turísticos validando usuario y lugar.
// @host localhost:5005
// @BasePath /api

// Inicia el servidor
func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RegisterReviewRoutes(r)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Middleware CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Iniciar el servidor
	log.Println("✅ Servidor corriendo en http://localhost:5005")
	log.Fatal(http.ListenAndServe(":5005", corsHandler.Handler(r)))
}
