package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRequest request params
type GetRequest struct {
	// API version
	Version string `json:"v"`
	Id      string `form:"user_id"`
}

// GetResponse return user
type GetResponse struct {
	// API version
	Version string      `json:"v"`
	User    models.User `json:"user"`
}

// GetUser godoc
// @Summary Return user by id
// @Description Return info about user by id
// @ID get-user
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param user_id query string true "user id"
// @Success 200 {object} handlers.GetResponse
// @Failure 400 {object} handlers.ResponseData
// @Router /user.get [Get]
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
