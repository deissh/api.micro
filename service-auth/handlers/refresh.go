package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
// @Summary Deactivate old token and create new
// @Description Generate new access_token and refresh_token
// @ID refresh-token
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param refresh_token query string false "refresh_token"
// @Success 200 {object} handlers.RefreshResponse
// @Failure 400 {object} handlers.ResponseData
// @Failure 500 {object} handlers.ResponseData
// @Router /token.refresh [Get]
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
	// remove token old token
	h.db.Delete(&token)

	var user models.User
	h.db.First(&user, token.UserID)

	jwttoken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := jwttoken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["role"] = user.Role
	claims["permissions"] = token.Permissions

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := jwttoken.SignedString([]byte(helpers.GetEnvWithPanic("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusInternalServerError,
			Data:   "JWT signing error",
		})
		return
	}

	refresh, err := helpers.GenerateRandomString(128)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseData{
			Status: http.StatusInternalServerError,
			Data:   "Refresh token generate error",
		})
		return
	}

	newToken := models.Token{
		AccessToken:  t,
		RefreshToken: refresh,
		UserID:       user.ID,
		Permissions:  token.Permissions,
	}

	h.db.Create(&newToken)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		Token:   newToken,
	})
}
