package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/service-auth/models"
	"net/http"
	"time"
)

// CheckRequest request params
type CheckRequest struct {
	// API version
	Version     string `form:"v"`
	AccessToken string `form:"access_token" binding:"required"`
}

type shortToken struct {
	UserID      uint     `json:"user_id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// CheckResponse default response
type CheckResponse struct {
	// API version
	Version string     `json:"v"`
	Token   shortToken `json:"token"`
}

// TokenCheck godoc
func (h Handler) TokenCheck(c *gin.Context) {
	var r CheckRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var token models.Token
	if err := h.db.Where(
		&models.Token{
			AccessToken: r.AccessToken,
		},
	).First(&token).Error; err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusUnauthorized,
			Data:   "Access token not founded",
		})
		return
	}

	if token.UpdatedAt.Add(time.Hour * 72).Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusUnauthorized,
			Data:   "Access token expired",
		})
		return
	}

	c.JSON(http.StatusOK, CheckResponse{
		Version: "1",
		Token: shortToken{
			UserID:      token.UserID,
			Role:        token.UserRole,
			Permissions: token.Permissions,
		},
	})
}
