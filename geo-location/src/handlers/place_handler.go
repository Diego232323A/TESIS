package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	config "geo-service/src/db"
	"geo-service/src/models"
)

// @Summary Agregar un nuevo lugar turístico
// @Description Agrega un nuevo lugar con nombre y coordenadas
// @Tags Places
// @Accept json
// @Produce json
// @Param place body models.Place true "Lugar turístico"
// @Success 201 {object} map[string]string
// @Failure 400 {string} string "Error al procesar la solicitud"
// @Router /places [post]

func CreatePlace(w http.ResponseWriter, r *http.Request) {
	var place models.Place
	err := json.NewDecoder(r.Body).Decode(&place)
	if err != nil {
		http.Error(w, "❌ Error en los datos de entrada", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO places (name, location) VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3), 4326)) RETURNING id`
	err = config.DB.QueryRow(query, place.Name, place.Lng, place.Lat).Scan(&place.ID)
	if err != nil {
		http.Error(w, "❌ Error al insertar en la base de datos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "📍 Lugar agregado correctamente"})
}

// @Summary Obtener lugares cercanos
// @Description Devuelve los lugares dentro de un radio de 5 km
// @Tags Places
// @Produce json
// @Param lat query float64 true "Latitud"
// @Param lng query float64 true "Longitud"
// @Success 200 {array} models.Place
// @Failure 400 {string} string "Error en los parámetros"
// @Router /places/nearby [get]
func GetNearbyPlaces(w http.ResponseWriter, r *http.Request) {
	lat, err1 := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lng, err2 := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	if err1 != nil || err2 != nil {
		http.Error(w, "❌ Parámetros inválidos", http.StatusBadRequest)
		return
	}

	query := `SELECT id, name, ST_X(location) AS lng, ST_Y(location) AS lat 
	          FROM places WHERE ST_DWithin(location, ST_SetSRID(ST_MakePoint($1, $2), 4326), 5000)`

	rows, err := config.DB.Query(query, lng, lat)
	if err != nil {
		http.Error(w, "❌ Error al buscar lugares cercanos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var places []models.Place
	for rows.Next() {
		var place models.Place
		if err := rows.Scan(&place.ID, &place.Name, &place.Lng, &place.Lat); err != nil {
			http.Error(w, "❌ Error al procesar los datos", http.StatusInternalServerError)
			return
		}
		places = append(places, place)
	}

	json.NewEncoder(w).Encode(places)
}
