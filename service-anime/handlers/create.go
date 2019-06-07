package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
	Strings "strings"
)

type anime struct {
	Title       string              `json:"title" binding:"required"`
	TitleEn     string              `json:"title_en" binding:"required"`
	TitleOr     string              `json:"title_or" binding:"required"`
	Annotation  string              `json:"annotation" binding:"required"`
	Description string              `json:"description" binding:"required"`
	Posters     string              `json:"posters" binding:"required"`
	Type        string              `json:"type" binding:"required"`
	Genres      string              `json:"genres" binding:"required"`
	Status      string              `json:"status" binding:"required"`
	Year        string              `json:"year" binding:"required"`
	Translators []models.Translator `json:"translators"`
	WorldArtID  string              `json:"world_art_id"`
	KinopoiskID string              `json:"kinopoisk_id"`
	Countries   string              `json:"countries"`
	Actors      string              `json:"actors"`
	Directors   string              `json:"directors"`
	Studios     string              `json:"studios"`
}

// CreateRequest request params
type CreateRequest struct {
	Anime anime `json:"anime"`
}

// CreateResponse return struct in response
type CreateResponse struct {
	// API version
	Version string       `json:"v"`
	Anime   models.Anime `json:"anime"`
}

// CreateAnime godoc
// @Summary Create anime
// @Description Create anime and return it
// @ID create-anime
// @Accept  json
// @Produce  json
// @Param anime body handlers.anime true "anime body"
// @Param v query string false "service version"
// @Param access_token query string true "user access_token"
// @Success 200 {object} handlers.CreateResponse
// @Failure 400 {object} handlers.ResponseData
// @Router /anime.create [Post]
func (h Handler) CreateAnime(c *gin.Context) {
	var r CreateRequest
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
		Title:       r.Anime.Title,
		TitleEn:     r.Anime.TitleEn,
		TitleOr:     r.Anime.TitleOr,
		Year:        r.Anime.Year,
		Genres:      Strings.Split(r.Anime.Genres, ","),
		Posters:     Strings.Split(r.Anime.Posters, ","),
		Annotation:  r.Anime.Annotation,
		Description: r.Anime.Description,
		Status:      r.Anime.Status,
		Type:        r.Anime.Type,
		KinopoiskID: r.Anime.KinopoiskID,
		WorldArtID:  r.Anime.WorldArtID,
		Translators: r.Anime.Translators,
		Countries:   Strings.Split(r.Anime.Countries, ","),
		Actors:      Strings.Split(r.Anime.Actors, ","),
		Directors:   Strings.Split(r.Anime.Directors, ","),
		Studios:     Strings.Split(r.Anime.Studios, ","),
	}

	h.db.Create(&anime)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		Anime:   anime,
	})
}
