package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"net/http"
)

// UpdateRequest request params
type UpdateRequest struct {
	News news `json:"news"`
}

// UpdateResponse response struct
type UpdateResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

// UpdateNews godoc
// @Summary update news
// @Description Update news and return it
// @ID update-news
// @Accept  json
// @Produce  json
// @Param news_id query string false "news_id"
// @Param news body handlers.news true "news body"
// @Param v query query false "service version"
// @Param access_token query string true "user access_token"
// @Success 200 {object} handlers.UpdateRequest
// @Failure 400 {object} handlers.ResponseData
// @Router /news.create [Post]
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
		c.DefaultQuery("access_token", ""),
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
	if err := h.db.First(&news, c.DefaultQuery("news_id", "")).Error; err != nil {
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

	// merge two struct
	if err := mergo.Merge(&r, news); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	// merge two slices to r
	newNews := models.News{
		Title:      r.News.Title,
		Annotation: r.News.Annotation,
		Body:       r.News.Body,
		Preview:    r.News.Preview,
		Background: r.News.Background,
		Types:      r.News.Types,
	}

	h.db.Create(&newNews)

	c.JSON(http.StatusOK, UpdateResponse{
		Version: "1",
		News:    newNews,
	})
}
