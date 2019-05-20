package handlers

import (
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRequest request params
type GetRequest struct {
	Id string `form:"news_id"`
}

// GetResponse response struct
type GetResponse struct {
	// API version
	Version string      `json:"v"`
	News    models.News `json:"news"`
}

// GetNews godoc
// @Summary Return news by id
// @Description Return info about news by id
// @ID get-news
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param news query string true "news id"
// @Success 200 {object} handlers.GetResponse
// @Failure 400 {object} handlers.ResponseData
// @Router /news.get [Get]
func (h Handler) GetNews(c *gin.Context) {

	var r GetRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	var news models.News
	if err := h.db.First(&news, r.Id).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Bad title name",
		})
		return
	}

	c.JSON(http.StatusOK, GetResponse{
		Version: "1",
		News:    news,
	})
}
