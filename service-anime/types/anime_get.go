package types

import "github.com/nekko-ru/api/models"

// GetRequest request params
type GetRequest struct {
	ID int `form:"anime_id"`
}

// GetResponse response struct
type GetResponse struct {
	// API version
	Version string       `json:"v"`
	Anime   models.Anime `json:"anime"`
}
