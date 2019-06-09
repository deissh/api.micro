package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckRequest request params
type ActivateRequest struct {
	// API version
	Version string `form:"v"`
	Token   string `form:"token" binding:"required"`
}

// TokenCheck godoc
func (h Handler) Activate(c *gin.Context) {
	var r ActivateRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var token models.ActivateTokens
	if err := h.db.Where(
		&models.ActivateTokens{
			Token: r.Token,
		},
	).First(&token).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "undefended activate token",
		})
		return
	}
	h.db.Delete(&token)
	h.db.LogMode(true)
	h.db.Exec("UPDATE \"users\" SET \"activated\" = true WHERE \"email\" = ?", token.Email)

	c.JSON(http.StatusOK, gin.H{})
}
