package models

type TMBDMovieResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"original_title"`
	Overview string `json:"overview"`
}
