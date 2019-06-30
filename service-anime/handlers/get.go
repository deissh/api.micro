package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRequest request params
type GetRequest struct {
	ID string `form:"anime_id"`
}

// GetResponse response struct
type GetResponse struct {
	// API version
	Version string       `json:"v"`
	Anime   models.Anime `json:"anime"`
}

// GetAnime godoc
func (h Handler) GetAnime(c *gin.Context) {

	var r GetRequest

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
			Data:   "News does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		Version: "1",
		Anime:   anime,
	})
}
