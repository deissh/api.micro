package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Summary Deactivate old token and create new
// @Description Check access_token
// @ID refresh-token
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param access_token query string false "access_token"
// @Success 200 {object} handlers.CheckResponse
// @Failure 400 {object} handlers.ResponseData
// @Failure 500 {object} handlers.ResponseData
// @Router /token.check [Get]
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
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Access token not founded",
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