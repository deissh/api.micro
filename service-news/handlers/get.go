package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetRequest struct {
	Id string `form:"news_id"`
}

type GetResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

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
	if err := h.db.First(&news, r.Id).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad title name",
		})
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		Version: "1",
		News:    news,
	})
}
