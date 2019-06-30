package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateRequest params
type UpdateRequest struct {
	// API version
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Picture   string `json:"picture"`
	Desc      string `json:"desc"`
	Status    string `json:"status"`
	Sex       int    `json:"sex"` // 1 – female; 2 – male.
}

// UpdateResponse response struct
type UpdateResponse struct {
	// API version
	Version string      `json:"v"`
	User    models.User `json:"user"`
}

// UpdateProfile update user info
func (h Handler) UpdateProfile(c *gin.Context) {
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

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

	var currentUser models.User
	if err := h.db.First(&currentUser, token.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "User does not exist",
		})
		return
	}

	h.db.Model(&currentUser).Update(r)

	c.JSON(http.StatusOK, UpdateResponse{
		Version: "1",
		User:    currentUser,
	})
}
