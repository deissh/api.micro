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
	ID      string `form:"news_id" binding:"required"`

	AccessToken string `form:"access_token" binding:"required"`
}

// RemoveResponse response struct
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

	_, err := helpers.TokenVerify(
		r.AccessToken,
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

	if err := h.db.Delete(&models.News{}, r.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "News not founded",
		})
		return
	}

	c.JSON(http.StatusOK, RemoveResponse{
		Version: "1",
		Status:  "deleted",
	})
}
