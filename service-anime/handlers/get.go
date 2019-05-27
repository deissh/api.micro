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
	Version string               `json:"v"`
	Anime   models.AnimeMoonWalk `json:"anime"`
}

// GetAnime godoc
// @Summary Return anime by id
// @Description Return info about anime by id
// @ID get-anime
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param anime query string true "anime id"
// @Success 200 {object} handlers.GetResponse
// @Failure 400 {object} handlers.ResponseData
// @Router /anime.get [Get]
func (h Handler) GetAnime(c *gin.Context) {

	var r GetRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var anime models.AnimeMoonWalk
	if err := h.db.First(&anime, r.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad title name",
		})
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		Version: "1",
		Anime:   anime,
	})
}
