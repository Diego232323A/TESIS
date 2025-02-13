package handlers

import (
	"database/sql"
	"delete-review/src/config"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Eliminar reseña
// @Description Permite eliminar una reseña si el usuario es el propietario.
// @Tags Reseñas
// @Param id path int true "ID de la reseña"
// @Param user_id query string true "ID del usuario propietario"
// @Success 200 {object} map[string]string "Reseña eliminada correctamente"
// @Failure 400 {object} map[string]string "Error en la solicitud"
// @Failure 404 {object} map[string]string "Reseña no encontrada"
// @Failure 403 {object} map[string]string "Usuario no autorizado"
// @Router /reviews/{id} [delete]
func DeleteReview(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	// Obtener el ID de la reseña desde la URL
	vars := mux.Vars(r)
	reviewID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "ID inválido"}`, http.StatusBadRequest)
		return
	}

	// Obtener el user_id de la consulta
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, `{"error": "Falta el user_id"}`, http.StatusBadRequest)
		return
	}

	// Verificar si la reseña existe y obtener su propietario
	var reviewOwnerID string
	err = db.QueryRow("SELECT user_id FROM reviews WHERE id=$1", reviewID).Scan(&reviewOwnerID)
	if err == sql.ErrNoRows {
		http.Error(w, `{"error": "Reseña no encontrada"}`, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, `{"error": "Error en la base de datos"}`, http.StatusInternalServerError)
		return
	}

	// Verificar que el usuario es el propietario de la reseña
	if reviewOwnerID != userID {
		http.Error(w, `{"error": "Usuario no autorizado para eliminar esta reseña"}`, http.StatusForbidden)
		return
	}

	// Eliminar la reseña
	_, err = db.Exec("DELETE FROM reviews WHERE id=$1", reviewID)
	if err != nil {
		http.Error(w, `{"error": "Error al eliminar la reseña"}`, http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Reseña eliminada correctamente"}`))
}
