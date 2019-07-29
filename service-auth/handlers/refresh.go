package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/service-auth/helpers"
	"github.com/nekko-ru/api/service-auth/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// RefreshRequest request params
type RefreshRequest struct {
	// API version
	Version      string `form:"v"`
	RefreshToken string `form:"refresh_token" binding:"required"`
}

// RefreshResponse response struct
type RefreshResponse struct {
	// API version
	Version string       `json:"v"`
	Token   models.Token `json:"token"`
}

// TokenRefresh godoc
func (h Handler) TokenRefresh(c *gin.Context) {
	var r RefreshRequest
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
			RefreshToken: r.RefreshToken,
		},
	).First(&token).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Access token not founded",
		})
		return
	}

	access, err := helpers.GenerateRandomString(128)
	if err != nil {
		log.Error("Access token generate error")
		c.JSON(http.StatusInternalServerError, ResponseData{
			Status: http.StatusInternalServerError,
			Data:   "Access token generate error",
		})
		return
	}

	refresh, err := helpers.GenerateRandomString(254)
	if err != nil {
		log.Error("Refresh token generate error")
		c.JSON(http.StatusInternalServerError, ResponseData{
			Status: http.StatusInternalServerError,
			Data:   "Refresh token generate error",
		})
		return
	}

	// update access_token and refresh
	token.AccessToken = refresh
	token.AccessToken = access

	h.db.Save(&token)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		Token:   token,
	})
}
