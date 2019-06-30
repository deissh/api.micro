package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRequest request params
type GetRequest struct {
	// API version
	Version string `json:"v"`
	ID      string `form:"user_id"`
}

// GetResponse return user
type GetResponse struct {
	// API version
	Version string      `json:"v"`
	User    models.User `json:"user"`
}

// GetUser godoc
func (h Handler) GetUser(c *gin.Context) {

	var r GetRequest
	var user models.User

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	if err := h.db.First(&user, r.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad nickname",
		})
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		User: user,
	})
}
