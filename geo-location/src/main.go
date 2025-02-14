package main

import (
	"fmt"
	"log"
	"net/http"

	config "geo-service/src/db"
	"geo-service/src/routes"

	"github.com/gorilla/mux"

	_ "geo-service/src/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.ConnectDB()

	router := mux.NewRouter()
	routes.SetupRoutes(router)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("🚀 Servidor ejecutándose en el puerto 5012...")
	log.Fatal(http.ListenAndServe(":5012", router))
}
