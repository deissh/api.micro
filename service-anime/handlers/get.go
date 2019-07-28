package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/service-anime/types"
	"net/http"
)

// GetAnime godoc
func (h Handler) GetAnime(c *gin.Context) {

	var r types.GetRequest

	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	anime, err := h.srv.Get(r.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, types.GetResponse{
		Version: "1",
		Anime:   anime,
	})
}
