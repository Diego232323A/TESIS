package models

// Review representa la estructura de una reseña en la BD
type Review struct {
	ID         int    `json:"id"`
	UserID     string `json:"user_id"`
	Place_name string `json: place_name`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
}
