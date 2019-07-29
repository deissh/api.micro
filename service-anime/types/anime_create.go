package types

import "github.com/nekko-ru/api/service-anime/models"

// CreateRequest contain necessary params
type CreateRequest struct {
	Title       string              `json:"title" binding:"required"`
	TitleEn     string              `json:"title_en" binding:"required"`
	TitleOr     string              `json:"title_or" binding:"required"`
	Annotation  string              `json:"annotation" binding:"required"`
	Description string              `json:"description" binding:"required"`
	Posters     []string            `json:"posters" binding:"required"`
	Type        string              `json:"type" binding:"required"`
	Genres      []string            `json:"genres" binding:"required"`
	Translators []models.Translator `gorm:"foreignkey:ID" json:"translators"`
	Status      string              `json:"status" binding:"required"`
	Year        string              `json:"year" binding:"required"`
	WorldArtID  string              `json:"world_art_id"`
	KinopoiskID string              `json:"kinopoisk_id"`
	Rating      float32             `json:"rating"`
	Votes       int                 `json:"votes"`
	Countries   []string            `json:"countries"`
	Actors      []string            `json:"actors"`
	Directors   []string            `json:"directors"`
	Studios     []string            `json:"studios"`
}

// CreateResponse return struct in response
type CreateResponse struct {
	// API version
	Version string       `json:"v"`
	Anime   models.Anime `json:"anime"`
}
