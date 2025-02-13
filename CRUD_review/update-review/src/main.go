package main

import (
	"log"
	"net/http"

	"update-review/src/config"
	"update-review/src/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "update-review/src/docs"
)

// @title Microservicio de Actualización de Reseñas
// @version 1.0
// @description API para actualizar reseñas de lugares turísticos

func main() {
	// Conectar a la BD
	config.ConnectDB()

	// Configurar router
	r := mux.NewRouter()
	routes.RegisterReviewRoutes(r)

	// Swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Middleware CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Iniciar servidor
	log.Println("Servidor corriendo en http://localhost:5006")
	log.Fatal(http.ListenAndServe(":5006", corsHandler.Handler(r)))
}
