package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/models"
	"net/http"
)

// RemoveRequest request params
type RemoveRequest struct {
	// API version
	Version     string `json:"v" query:"v"`
	AccessToken string `form:"access_token" binding:"required"`
	All         bool   `form:"all"`
}

// RemoveResponse response struct
type RemoveResponse struct {
	// API version
	Version string `json:"v"`
	Status  string `json:"status"`
}

// TokenRemove godoc
func (h Handler) TokenRemove(c *gin.Context) {
	var r RemoveRequest
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
	// remove token
	h.db.Delete(&token)

	if r.All == true {
		h.db.Where(&models.Token{
			UserID: token.UserID,
		}).Delete(models.Token{})
	}

	c.JSON(http.StatusOK, RemoveResponse{
		Version: "1",
		Status:  "ok",
	})
}
