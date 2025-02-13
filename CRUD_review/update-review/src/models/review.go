package models

// Review representa una reseña turística
// @Description Modelo de actualización de reseña
type Review struct {
	ID      int    `json:"id"` // ID de la reseña
	User_id int    `json:"user_id"`
	Rating  int    `json:"rating"`  // Nueva puntuación del lugar
	Comment string `json:"comment"` // Nuevo comentario
}
