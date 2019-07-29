package types

import "github.com/nekko-ru/api/service-anime/models"

// FindRequest request params
type FindRequest struct {
	Query string `form:"q"`
	//Genres []string `form:"genres"`
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Sort  string `form:"sort"`
}

// FindResponse response struct
type FindResponse struct {
	// API version
	Version string              `json:"v"`
	Animes  []models.AnimeShort `json:"animes"`
}
