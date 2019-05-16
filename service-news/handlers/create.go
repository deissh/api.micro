package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateNewsR struct {
	Title      string `form:"title" binding:"required"`
	Annotation string `form:"annotation" binding:"required"`
	Body       string `form:"body" binding:"required"`
	Preview    string `form:"preview" binding:"required"`
	Background string `form:"background"`
	Types      string `form:"types"`

	AccessToken string `form:"access_token" binding:"required"`
}

type CreateResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

func (h Handler) CreateNews(c *gin.Context) {
	var r CreateNewsR
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	token, err := helpers.TokenVerify(
		r.AccessToken,
		true,
		[]string{"newsmaker", "admin", "superadmin"},
		[]string{"news"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Unauthorized",
		})
		return
	}

	var author models.User

	if err := h.db.First(&author, token.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad auth",
		})
		return
	}

	news := models.News{
		Title:      r.Title,
		Annotation: r.Annotation,
		Body:       r.Body,
		Author:     author,
		Preview:    r.Preview,
		Background: r.Background,
		Types:      r.Types,
	}

	h.db.Create(&news)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		News:    news,
	})
}
