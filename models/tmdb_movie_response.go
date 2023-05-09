package models

type TMBDMovieResponse struct {
	ID           int     `json:"id"`
	Title        string  `json:"original_title"`
	Overview     string  `json:"overview"`
	Popularity   float64 `json:"popularity"`
	Release_Date string  `json:"release_date"`
	Rating       float64 `json:"vote_average"`
}

type TrendingMoviesResponse struct {
	Results []TMBDMovieResponse `json:"results"`
}
