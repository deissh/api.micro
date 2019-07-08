package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/models"
	"net/http"
)

// PasswordRestoreRequest godoc
type PasswordRestoreRequest struct {
	Email string `json:"email" binding:"required"`
}

// PasswordRestore godoc
func (h Handler) PasswordRestore(c *gin.Context) {
	var r PasswordRestoreRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var user models.User
	if err := h.db.First(&user, models.User{Email: r.Email}).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "user not founded",
		})
		return
	}

	// sending activation email
	activateToken, err := helpers.GenerateRandomString(32)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "server error",
		})
		return
	}
	h.db.Create(&models.PasswordRestoreTokens{
		Email: user.Email,
		Token: activateToken,
	})
	_ = helpers.SendEmail(
		helpers.PasswordRestoreTemplate,
		user.Email,
		map[string]string{
			"restore_url": "https://anibe.ru/account/restore?token=" + activateToken,
			"first_name":  user.FirstName,
			"last_name":   user.LastName,
		},
	)

	c.JSON(http.StatusOK, ResponseData{
		Status: http.StatusOK,
		Data:   "email with restore link has been send",
	})
}
