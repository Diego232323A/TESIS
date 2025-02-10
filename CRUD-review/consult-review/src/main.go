package main

import (
	"consult-review/src/config"
	"consult-review/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "consult-review/src/docs"
)

// @title Microservicio de Consulta de Reseñas
// @version 1.0
// @description API para obtener reseñas de lugares turísticos.

func main() {
	// Conectar a la BD
	config.ConnectDB()

	// Configurar router
	r := mux.NewRouter()

	// Configurar rutas
	routes.RegisterReviewRoutes(r)

	// Swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Middleware CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Iniciar servidor
	log.Println("Servidor corriendo en http://localhost:5008")
	log.Fatal(http.ListenAndServe(":5008", corsHandler.Handler(r)))
}
