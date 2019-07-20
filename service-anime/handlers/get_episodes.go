package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/models"
	"net/http"
)

// GetRequest request params
type GetEpisodesRequest struct {
	AnimeID      uint `form:"anime_id"`
	TranslatorID uint `form:"translator_id"`
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

	var transl models.Translator
	if err := h.db.Where(&models.Translator{ID: r.TranslatorID, AnimeID: r.AnimeID}).First(&transl).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Anime does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, GetEpisodesResponse{
		Version:  "1",
		Episodes: transl.Episodes,
	})
}
