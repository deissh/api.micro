package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/service-anime/helpers"
	"github.com/nekko-ru/api/service-anime/types"
	"net/http"
)

// CreateAnime godoc
func (h Handler) CreateAnime(c *gin.Context) {
	var r types.CreateRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error: " + err.Error(),
		})
		return
	}

	_, err := helpers.TokenVerify(
		c.DefaultQuery("access_token", ""),
		true,
		[]string{"moderator", "admin", "superadmin"},
		[]string{"anime"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusUnauthorized,
			Data:   err.Error(),
		})
		return
	}

	anime, err := h.Srv.Create(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.CreateResponse{
		Version: "1",
		Anime:   anime,
	})
}
