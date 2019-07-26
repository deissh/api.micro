package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/models"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// CreateRequest request params
type CreateRequest struct {
	// API version
	Version  string `json:"v" query:"v"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Scope    string `form:"scope" binding:"required"`
}

// CreateResponse response struct
type CreateResponse struct {
	// API version
	Version string       `json:"v"`
	Token   models.Token `json:"token"`
}

// TokenCreate godoc
func (h Handler) TokenCreate(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var user models.User
	if err := h.db.Where(&models.User{Email: r.Email, Activated: true}).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad password or email",
		})
		return
	}
	// сделал так как используется BCrypt на строне сервера
	if err := user.CheckPassword(r.Password); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad password or email",
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

	refresh, err := helpers.GenerateRandomString(128)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseData{
			Status: http.StatusInternalServerError,
			Data:   "Refresh token generate error",
		})
		return
	}

	token := models.Token{
		AccessToken:  access,
		RefreshToken: refresh,
		UserID:       user.ID,
		UserRole:     user.Role,
		Permissions:  strings.Split(r.Scope, ","),
	}

	h.db.Create(&token)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		Token:   token,
	})
}
