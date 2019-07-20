package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/models"
	"net/http"
)

// GetRequest request params
type GetEpisodesRequest struct {
	ID           string `form:"anime_id"`
	TranslatorID uint   `form:"translator_id"`
}

// GetResponse response struct
type GetEpisodesResponse struct {
	// API version
	Version  string   `json:"v"`
	Episodes []string `json:"episodes"`
}

// GetAnime godoc
func (h Handler) GetEpisodesAnime(c *gin.Context) {
	var r GetEpisodesRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var anime models.Anime
	if err := h.db.First(&anime, r.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Anime does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, GetEpisodesResponse{
		Version:  "1",
		Episodes: anime.GetEpisodesByTranslator(r.TranslatorID),
	})
}
