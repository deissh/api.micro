package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateNewsR request params
type CreateNewsR struct {
	Title      string `form:"title" binding:"required"`
	Annotation string `form:"annotation" binding:"required"`
	Body       string `form:"body" binding:"required"`
	Preview    string `form:"preview" binding:"required"`
	Background string `form:"background"`
	Types      string `form:"types"`

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
// @Param v query string false "service version"
// @Param title query string true "title"
// @Param annotation query string true "annotation"
// @Param body query string true "body news"
// @Param preview query string true "preview"
// @Param background query string false "background"
// @Param types query string false "news types"
// @Param access_token query string true "user access_token"
// @Success 200 {object} handlers.CreateResponse
// @Failure 400 {object} handlers.ResponseData
// @Router /news.create [Get]
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
