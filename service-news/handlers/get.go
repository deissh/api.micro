package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/models"
	"net/http"
)

// GetRequest request params
type GetRequest struct {
	ID string `form:"news_id"`
}

// GetResponse response struct
type GetResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

// GetNews godoc
func (h Handler) GetNews(c *gin.Context) {

	var r GetRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var news models.News
	if err := h.db.First(&news, r.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "News does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		Version: "1",
		News:    news,
	})
}
