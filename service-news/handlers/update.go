package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateRequest request params
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

// UpdateResponse response struct
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

// UpdateNews godoc
// @Summary update news
// @Description Update news and return it
// @ID update-news
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param title query string false "title"
// @Param annotation query string false "annotation"
// @Param body query string false "body news"
// @Param preview query string false "preview"
// @Param background query string false "background"
// @Param types query string false "news types"
// @Param access_token query string true "user access_token"
// @Success 200 {object} handlers.UpdateRequest
// @Failure 400 {object} handlers.ResponseData
// @Router /news.create [Get]
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
