package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RemoveRequest request params
type RemoveRequest struct {
	// API version
	Version string `json:"v" query:"v"`
	ID      string `form:"anime_id" binding:"required"`

	AccessToken string `form:"access_token" binding:"required"`
}

// RemoveResponse response struct
type RemoveResponse struct {
	// API version
	Version string `json:"v"`
	Status  string `json:"status"`
}

// RemoveAnime godoc
// @Summary Remove anime by id
// @Description Remove anime by id
// @ID remove-anime
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param anime_id query string true "anime id"
// @Success 200 {object} handlers.RemoveResponse
// @Failure 400 {object} handlers.ResponseData
// @Router /anime.remove [Get]
func (h Handler) RemoveAnime(c *gin.Context) {
	var r RemoveRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	_, err := helpers.TokenVerify(
		r.AccessToken,
		true,
		[]string{"moderator", "admin", "superadmin"},
		[]string{"anime"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Unauthorized",
		})
		return
	}

	if err := h.db.Delete(&models.Anime{}, r.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "AnimeMoonWalk not founded",
		})
		return
	}

	c.JSON(http.StatusOK, RemoveResponse{
		Version: "1",
		Status:  "deleted",
	})
}
