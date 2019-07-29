package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/service-anime/helpers"
	"github.com/nekko-ru/api/service-anime/models"
	"net/http"
	"strconv"
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

	if _, err := helpers.TokenVerify(
		c.DefaultQuery("access_token", ""),
		true,
		[]string{"animemaker", "admin", "superadmin"},
		[]string{"anime"},
	); err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Unauthorized",
		})
		return
	}

	id, err := strconv.Atoi(c.DefaultQuery("anime_id", ""))
	anime, err := h.Srv.Update(id, r)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Error in params",
		})
		return
	}

	c.JSON(http.StatusOK, UpdateResponse{
		Version: "1",
		Anime:   anime,
	})
}
