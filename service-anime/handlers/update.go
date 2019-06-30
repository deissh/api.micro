package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"net/http"
)

// UpdateResponse response struct
type UpdateResponse struct {
	// API version
	Version string       `json:"v"`
	Anime   models.Anime `json:"anime"`
}

// UpdateAnime godoc
func (h Handler) UpdateAnime(c *gin.Context) {
	var r models.Anime
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	_, err := helpers.TokenVerify(
		c.DefaultQuery("access_token", ""),
		true,
		[]string{"animemaker", "admin", "superadmin"},
		[]string{"anime"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Unauthorized",
		})
		return
	}

	var anime models.Anime
	if err := h.db.First(&anime, c.DefaultQuery("anime_id", "")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "News does not exist",
		})
		return
	}

	// merge two struct
	if err := mergo.Merge(&r, anime); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	h.db.Model(&anime).Update(r)

	c.JSON(http.StatusOK, UpdateResponse{
		Version: "1",
		Anime:   anime,
	})
}
