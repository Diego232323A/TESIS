package models

// Review representa una reseña de un lugar turístico
// @Description Modelo de reseña
type Review struct {
	ID       int    `json:"id"`
	User_id  int    `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating"`
}
