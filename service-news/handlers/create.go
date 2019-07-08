package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/models"
	"net/http"
)

type news struct {
	Title      string `json:"title" binding:"required"`
	Annotation string `json:"annotation" binding:"required"`
	Body       string `json:"body" binding:"required"`
	Preview    string `json:"preview" binding:"required"`
	Background string `json:"background"`
	Types      string `json:"types"`
}

// CreateResponse return struct in response
type CreateResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

// CreateNews godoc
func (h Handler) CreateNews(c *gin.Context) {
	var r news
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
		[]string{"newsmaker", "admin", "superadmin"},
		[]string{"news"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusBadRequest,
			Data:   err.Error(),
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
