package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/models"
	"net/http"
)

// AnimeParams contain necessary params
type AnimeParams struct {
	Title       string   `json:"title" binding:"required"`
	TitleEn     string   `json:"title_en" binding:"required"`
	TitleOr     string   `json:"title_or" binding:"required"`
	Annotation  string   `json:"annotation" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Posters     []string `json:"posters" binding:"required"`
	Type        string   `json:"type" binding:"required"`
	Genres      []string `json:"genres" binding:"required"`
	Status      string   `json:"status" binding:"required"`
	Year        string   `json:"year" binding:"required"`
	WorldArtID  string   `json:"world_art_id"`
	KinopoiskID string   `json:"kinopoisk_id"`
	Countries   []string `json:"countries"`
	Actors      []string `json:"actors"`
	Directors   []string `json:"directors"`
	Studios     []string `json:"studios"`
}

// CreateResponse return struct in response
type CreateResponse struct {
	// API version
	Version string            `json:"v"`
	Anime   models.AnimeShort `json:"anime"`
}

// CreateAnime godoc
func (h Handler) CreateAnime(c *gin.Context) {
	var r AnimeParams
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error: " + err.Error(),
		})
		return
	}

	token, err := helpers.TokenVerify(
		c.DefaultQuery("access_token", ""),
		true,
		[]string{"moderator", "admin", "superadmin"},
		[]string{"anime"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusUnauthorized,
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

	anime := models.Anime{
		Title:       r.Title,
		TitleEn:     r.TitleEn,
		TitleOr:     r.TitleOr,
		Year:        r.Year,
		Genres:      r.Genres,
		Posters:     r.Posters,
		Annotation:  r.Annotation,
		Description: r.Description,
		Status:      r.Status,
		Type:        r.Type,
		KinopoiskID: r.KinopoiskID,
		WorldArtID:  r.WorldArtID,
		Countries:   r.Countries,
		Actors:      r.Actors,
		Directors:   r.Directors,
		Studios:     r.Studios,
	}

	h.db.Create(&anime)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		Anime:   anime.ViewShort(),
	})
}
