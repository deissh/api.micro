package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/models"
	"net/http"
)

// GetResponse return user
type GetResponse struct {
	// API version
	Version string      `json:"v"`
	User    models.User `json:"user"`
}

// GetProfile godoc
func (h Handler) GetProfile(c *gin.Context) {
	var user models.User

	token, err := helpers.TokenVerify(
		c.DefaultQuery("access_token", ""),
		true,
		// allow to all roles (but not banned roles)
		[]string{"newsmaker", "developer", "moderator", "admin", "superadmin", "user", "supporter"},
		[]string{"account"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusUnauthorized,
			Data:   err.Error(),
		})
		return
	}

	if err := h.db.First(&user, token.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "User not founded",
		})
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		User: user,
	})
}
