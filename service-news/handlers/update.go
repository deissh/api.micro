package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/models"
	"net/http"
)

// UpdateResponse response struct
type UpdateResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

// UpdateNews godoc
func (h Handler) UpdateNews(c *gin.Context) {
	var r models.News
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	token, err := helpers.TokenVerify(
		c.DefaultQuery("access_token", ""),
		true,
		[]string{"newsmaker", "admin", "superadmin"},
		[]string{"news"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Unauthorized",
		})
		return
	}

	var news models.News
	if err := h.db.First(&news, c.DefaultQuery("news_id", "")).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "News did not find",
		})
		return
	}

	var author models.User
	if err := h.db.First(&author, token.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad auth",
		})
		return
	}

	// merge two struct
	if err := mergo.Merge(&r, news); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	r.Author = author
	h.db.Model(&news).Update(r)

	c.JSON(http.StatusOK, UpdateResponse{
		Version: "1",
		News:    news,
	})
}
