package handlers

import (
	"consult-review/src/config"
	"consult-review/src/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllReviews obtiene todas las reseñas
// @Summary Obtiene todas las reseñas
// @Description Devuelve un listado de todas las reseñas almacenadas en la base de datos.
// @Tags Reviews
// @Produce json
// @Success 200 {array} models.Review
// @Failure 500 {object} map[string]string
// @Router /reviews [get]
func GetAllReviews(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, user_id, place_name, rating, comment FROM reviews")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.Place_name, &review.Rating, &review.Comment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, review)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

// GetReviewByID obtiene una reseña por su ID
// @Summary Obtiene una reseña específica
// @Description Devuelve una reseña según el ID proporcionado en la URL.
// @Tags Reviews
// @Produce json
// @Param id path int true "ID de la reseña"
// @Success 200 {object} models.Review
// @Failure 404 {object} map[string]string
// @Router /reviews/{id} [get]
func GetReviewByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var review models.Review
	err := config.DB.QueryRow("SELECT id, user_id, place_name, rating, comment FROM reviews WHERE place_name = $1", id).
		Scan(&review.ID, &review.UserID, &review.Place_name, &review.Rating, &review.Comment)

	if err != nil {
		http.Error(w, "Reseña no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

// GetReviewsByPlace obtiene todas las reseñas de un lugar específico
// @Summary Obtiene reseñas de un lugar turístico
// @Description Devuelve todas las reseñas asociadas a un lugar, según el ID del lugar.
// @Tags Reviews
// @Produce json
// @Param place_id path string true "ID del lugar turístico"
// @Success 200 {array} models.Review
// @Failure 500 {object} map[string]string
// @Router /reviews/place/{place_id} [get]
func GetReviewsByPlace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	placeNAME := params["place_name"]

	rows, err := config.DB.Query("SELECT id, user_id, place_name, rating, comment FROM reviews WHERE place_name = $1", placeNAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.Place_name, &review.Rating, &review.Comment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, review)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}
