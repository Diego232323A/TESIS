package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// Importamos strconv para convertir int a string
	"create-review/src/config"
	"create-review/src/models"
)

func ValidateIDs(username, name string) (bool, error) {
	// Crear JSON de solicitud
	requestBody, err := json.Marshal(map[string]string{
		"valor_mongo":    name,     // name en MongoDB
		"valor_postgres": username, // username en PostgreSQL
	})
	if err != nil {
		return false, err
	}

	// Enviar la solicitud al microservicio de validación
	validationURL := os.Getenv("VALIDATION_SERVICE_URL") + "/buscar"
	resp, err := http.Post(validationURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Decodificar la respuesta
	var result map[string]bool
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	// Retornar el resultado
	return result["resultado"], nil
}

// @Summary Crear una reseña
// @Description Agrega una reseña validando que el usuario y el lugar existan.
// @Tags Reseñas
// @Accept json
// @Produce json
// @Param review body models.Review true "Datos de la reseña"
// @Success 201 {object} map[string]string "Reseña creada exitosamente"
// @Failure 400 {object} map[string]string "Error en la solicitud"
// @Failure 404 {object} map[string]string "Usuario o lugar no encontrado"
// @Router /reviews [post]
func CreateReviewHandler(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	validationURL := os.Getenv("VALIDATION_SERVICE_URL") + "/buscar"
	fmt.Println("📡 Enviando solicitud a:", validationURL) // Depuración

	// Validar usuario y lugar con el microservicio validation
	isValid, err := ValidateIDs(review.Username, review.Name)
	if err != nil {
		http.Error(w, "Error en la validación", http.StatusInternalServerError)
		return
	}

	if !isValid {
		http.Error(w, "Usuario o lugar no encontrado", http.StatusNotFound)
		return
	}

	pgConnStr := "host=localhost dbname=users_db user=postgres password=12345 sslmode=disable"
	pgDb, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer pgDb.Close()

	var dato string

	query := "SELECT id FROM users WHERE username = $1"
	err = pgDb.QueryRow(query, review.Username).Scan(&dato)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No se encontró el dato", http.StatusNotFound)
		} else {
			fmt.Println("Error en la consulta:", err) // Imprimir el error para más detalles
			http.Error(w, "Error en la consulta", http.StatusInternalServerError)
		}
		return
	}

	db := config.GetDB()
	_, err = db.Exec("INSERT INTO reviews (user_id, user_name, place_name, comment, rating) VALUES ($1, $2, $3, $4, $5)",
		dato, review.Username, review.Name, review.Comment, review.Rating)

	if err != nil {
		fmt.Println("Error en la consulta:", err) // Imprimir el error para más detalles
		http.Error(w, "Error al guardar la reseña", http.StatusInternalServerError)
		return
	}

	log.Println("✅ Reseña creada con éxito")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Reseña creada exitosamente"})
}
