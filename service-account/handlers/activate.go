package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ActivateRequest request params
type ActivateRequest struct {
	// API version
	Version string `form:"v"`
	Token   string `form:"token" binding:"required"`
}

// Activate godoc
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
	var us models.User
	h.db.Exec("UPDATE \"users\" SET \"activated\" = true WHERE \"email\" = ?", token.Email).First(&us)

	_ = helpers.SendEmail(
		helpers.ActivatedAccountTemplate,
		token.Email,
		map[string]string{
			"first_name": us.FirstName,
			"last_name":  us.LastName,
		},
	)

	c.JSON(http.StatusOK, ResponseData{
		Status: http.StatusOK,
		Data:   "activated",
	})
}
