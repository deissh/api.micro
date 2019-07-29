package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/service-account/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// PasswordChangeRequest request params
type PasswordChangeRequest struct {
	// API version
	Version  string `form:"v"`
	Token    string `json:"token" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// PasswordChange godoc
func (h Handler) PasswordChange(c *gin.Context) {
	var r PasswordChangeRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var token models.PasswordRestoreTokens
	if err := h.db.Where(
		&models.PasswordRestoreTokens{
			Token: r.Token,
		},
	).First(&token).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "undefended password change token",
		})
		return
	}
	h.db.Delete(&token)

	var us models.User
	h.db.First(&us)

	if err := us.SetPassword(r.Password); err != nil {
		log.Error(err)
	}

	c.JSON(http.StatusOK, ResponseData{
		Status: http.StatusOK,
		Data:   "changed",
	})
}
