package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/models"
	"net/http"
	"strings"
)

// FindRequest request params
type FindRequest struct {
	Query string `form:"q"`
	//Genres []string `form:"genres"`
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Sort  string `form:"sort"`
}

// FindResponse response struct
type FindResponse struct {
	// API version
	Version string              `json:"v"`
	Animes  []models.AnimeShort `json:"animes"`
}

// FindAnime godoc
func (h Handler) FindAnime(c *gin.Context) {

	var r FindRequest

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

	var animes []models.Anime
	var err error
	if len(r.Query) > 0 {
		err = h.db.Where(
			"to_tsvector(title) || to_tsvector(title_en) || to_tsvector(title_or) @@ plainto_tsquery(?)",
			r.Query,
		).Limit(r.Limit).Offset(r.Limit * r.Page).Order(strings.Replace(r.Sort, "-", " desc", 1)).Find(&animes).Error
	} else {
		err = h.db.Limit(r.Limit).Offset(r.Limit * r.Page).Order(strings.Replace(r.Sort, "-", " desc", 1)).Find(&animes).Error
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Anime does not exist",
		})
		return
	}

	var short_animes []models.AnimeShort
	for _, a := range animes {
		short_animes = append(short_animes, a.ViewShort())
	}

	c.JSON(http.StatusOK, FindResponse{
		Version: "1",
		Animes:  short_animes,
	})
}
