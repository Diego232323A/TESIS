package main

import (
	"fmt"
	"log"
	"net/http"

	_ "validation-user-place/docs" // Importa los documentos generados por swag
	"validation-user-place/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Mi API de Búsqueda
// @version 1.0
// @description Esta es una API para buscar en MongoDB y PostgreSQL.

// @host localhost:8080
// @BasePath /

// @contact.name Soporte Técnico
// @contact.url http://www.soporte.com
// @contact.email soporte@soporte.com

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	r := mux.NewRouter()
	r.HandleFunc("/buscar", handlers.Handler).Methods("POST")

	// Ruta para servir la documentación Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.Handle("/", r)
	fmt.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
