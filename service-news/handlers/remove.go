package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RemoveRequest struct {
	// API version
	Version string `json:"v" query:"v"`
	Id      string `form:"id" binding:"required"`
}

type RemoveResponse struct {
	// API version
	Version string `json:"v"`
	Status  string `json:"status"`
}

func (h Handler) RemoveNews(c *gin.Context) {
	var r RemoveRequest
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
			Data:   "News not founded",
		})
		return
	}
	h.db.Delete(&news)

	c.JSON(http.StatusOK, RemoveResponse{
		Version: "1",
		Status:  "deleted",
	})
}
