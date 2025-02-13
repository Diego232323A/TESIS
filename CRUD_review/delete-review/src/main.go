package main

import (
	"delete-review/src/config"
	"delete-review/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "delete-review/src/docs"
)

// @title Microservicio de Eliminación de Reseñas
// @version 1.0
// @description API para eliminar reseñas de lugares turísticos si el usuario es el propietario.

func main() {
	// Conectar a la BD
	config.ConnectDB()

	// Configurar router
	r := mux.NewRouter()

	// Registrar rutas
	routes.RegisterReviewRoutes(r)

	// Swagger docs
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Middleware CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Iniciar servidor
	log.Println("Servidor corriendo en http://localhost:5007")
	log.Fatal(http.ListenAndServe(":5007", corsHandler.Handler(r)))
}
