package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
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

// CreateRequest request params
type CreateRequest struct {
	News news `json:"news"`

	Version     string `form:"v"`
	AccessToken string `form:"access_token" binding:"required"`
}

// CreateResponse return struct in response
type CreateResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

// CreateNews godoc
// @Summary Create news
// @Description Create news and return it
// @ID create-news
// @Accept  json
// @Produce  json
// @Param news body handlers.news true "news body"
// @Param v query string false "service version"
// @Param access_token query string true "user access_token"
// @Success 200 {object} handlers.CreateResponse
// @Failure 400 {object} handlers.ResponseData
// @Router /news.create [Post]
func (h Handler) CreateNews(c *gin.Context) {
	var r CreateRequest
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
		Title:      r.News.Title,
		Annotation: r.News.Annotation,
		Body:       r.News.Body,
		Author:     author,
		Preview:    r.News.Preview,
		Background: r.News.Background,
		Types:      r.News.Types,
	}

	h.db.Create(&news)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		News:    news,
	})
}
