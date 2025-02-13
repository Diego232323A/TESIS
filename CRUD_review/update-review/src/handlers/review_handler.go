package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"update-review/src/config"
	"update-review/src/models"

	"github.com/gorilla/mux"
)

// @Summary Actualizar reseña
// @Description Permite actualizar una reseña existente, pero solo si el usuario es el propietario.
// @Tags Reseñas
// @Accept json
// @Produce json
// @Param id path int true "ID de la reseña"
// @Param review body models.Review true "Reseña actualizada"
// @Success 200 {object} map[string]string "Reseña actualizada correctamente"
// @Failure 400 {object} map[string]string "Error en la solicitud"
// @Failure 404 {object} map[string]string "Reseña no encontrada"
// @Failure 403 {object} map[string]string "Usuario no autorizado"
// @Router /reviews/{id} [put]
func UpdateReview(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Obtener el ID de la reseña desde la URL
	vars := mux.Vars(r)
	reviewID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "ID inválido"}`, http.StatusBadRequest)
		return
	}

	// Obtener el cuerpo de la solicitud (los nuevos datos de la reseña)
	var review models.Review
	err = json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		fmt.Println("Error en la consulta:", err)
		http.Error(w, `{"error": "Datos inválidos"}`, http.StatusBadRequest)
		return
	}

	// Aquí deberías obtener el ID del usuario (propietario de la reseña)
	// Suponiendo que lo pasas como parte de la solicitud o lo extraes del token de autenticación
	// Este es solo un ejemplo, ajusta según tu sistema de autenticación
	userID := review.User_id // Este sería el ID del usuario que está intentando actualizar la reseña

	// Verificar que el usuario es el propietario de la reseña
	var reviewOwnerID int
	err = db.QueryRow("SELECT user_id FROM reviews WHERE id=$1", reviewID).Scan(&reviewOwnerID)
	if err != nil {
		http.Error(w, `{"error": "Reseña no encontrada"}`, http.StatusNotFound)
		return
	}

	print(reviewOwnerID)

	if reviewOwnerID != userID {
		http.Error(w, `{"error": "Usuario no autorizado"}`, http.StatusForbidden)
		return
	}

	// Actualizar la reseña en la base de datos
	_, err = db.Exec(
		"UPDATE reviews SET rating=$1, comment=$2 WHERE id=$3",
		review.Rating,
		review.Comment,
		reviewID,
	)
	if err != nil {
		http.Error(w, `{"error": "Error al actualizar la reseña"}`, http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Reseña actualizada correctamente"}`))
}
