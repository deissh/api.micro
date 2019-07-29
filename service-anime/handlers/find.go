package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/service-anime/models"
	"github.com/nekko-ru/api/service-anime/types"
	"net/http"
)

// FindAnime godoc
func (h Handler) FindAnime(c *gin.Context) {

	var r types.FindRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	if r.Limit > 50 || r.Limit < 1 {
		r.Limit = 25
	}
	if r.Page < 0 || r.Page > 1000 {
		r.Page = 0
	}

	res, err := h.Srv.Search(r.Query, r.Limit, r.Page, r.Sort)
	if err != nil {
		c.JSON(http.StatusOK, types.FindResponse{
			Version: "1",
			Animes:  []models.AnimeShort{},
		})
		return
	}

	c.JSON(http.StatusOK, types.FindResponse{
		Version: "1",
		Animes:  res,
	})
}
