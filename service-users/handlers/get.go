package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetRequest struct {
	Id string `form:"user_id"`
}

type GetResponse struct {
	User models.User `json:"user"`
}

func (h Handler) GetUser(c *gin.Context) {

	var r GetRequest
	var user models.User

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	if err := h.db.First(&user, r.Id).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad nickname",
		})
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		User: user,
	})
}
