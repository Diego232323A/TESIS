package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"validation-user-place/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Request struct {
	ValorMongo    string `json:"valor_mongo"`
	ValorPostgres string `json:"valor_postgres"`
}

func buscarEnMongo(valor string, client *mongo.Client) bool {
	collection := client.Database("tourist_places").Collection("places")
	filter := bson.D{{"name", valor}}
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return err == nil
}

func buscarEnPostgres(valor string, db *sql.DB) bool {
	query := "SELECT username FROM users WHERE username = $1"
	row := db.QueryRow(query, valor)
	var resultado string
	err := row.Scan(&resultado)
	return err == nil
}

// Handler
// @Summary Busca en MongoDB y PostgreSQL
// @Description Este endpoint busca en ambas bases de datos y devuelve el resultado.
// @Accept  json
// @Produce  json
// @Param   request body Request true "Parámetros de búsqueda"
// @Success 200 {object} map[string]bool
// @Failure 400 {string} string "Bad Request"
// @Router /buscar [post]
func Handler(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Conexión a MongoDB
	mongoClient, mongoCtx, mongoCancel, err := db.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer mongoCancel()
	defer mongoClient.Disconnect(mongoCtx)

	// Conexión a PostgreSQL
	pgDb, err := db.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer pgDb.Close()

	// Realizar búsqueda en ambas bases de datos
	resultadoMongo := buscarEnMongo(req.ValorMongo, mongoClient)
	resultadoPostgres := buscarEnPostgres(req.ValorPostgres, pgDb)

	// Devolver verdadero si ambos resultados son verdaderos
	response := resultadoMongo && resultadoPostgres

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"resultado": response})
}
