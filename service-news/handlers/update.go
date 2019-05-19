package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateRequest struct {
	// API version
	Version    string `form:"v"`
	Id         string `form:"news_id"`
	Title      string `form:"title"`
	Annotation string `form:"annotation"`
	Body       string `form:"body"`
	Preview    string `form:"preview"`
	Background string `form:"background"`
	Types      string `form:"types"`

	AccessToken string `form:"access_token" binding:"required"`
}

type UpdateResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

func checkNull(old string, new string) string {
	if new == "" {
		new = old
	}
	return new
}

func (h Handler) UpdateNews(c *gin.Context) {
	var r UpdateRequest
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

	var news models.News
	if err := h.db.First(&news, r.Id).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "News did not find",
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

	newNews := models.News{
		Title:      checkNull(news.Title, r.Title),
		Annotation: checkNull(news.Annotation, r.Annotation),
		Body:       checkNull(news.Body, r.Body),
		Author:     author,
		Preview:    checkNull(news.Preview, r.Preview),
		Background: checkNull(news.Background, r.Background),
		Types:      checkNull(news.Types, r.Types),
	}

	h.db.Create(&newNews)

	c.JSON(http.StatusOK, UpdateResponse{
		Version: "1",
		News:    newNews,
	})
}
