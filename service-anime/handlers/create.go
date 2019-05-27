package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"net/http"
	Strings "strings"
)

type anime struct {
	TitleRu          string  `json:"title_ru"`
	TitleEn          string  `json:"title_en"`
	Year             int     `json:"year"`
	Genres           string  `json:"genres"`
	Poster           string  `json:"poster"`
	Tagline          string  `json:"tagline"`
	Description      string  `json:"description"`
	Token            string  `json:"token"`
	Type             string  `json:"type"`
	KinopoiskID      int     `json:"kinopoisk_id"`
	WorldArtID       int     `json:"world_art_id"`
	Translator       string  `json:"translator"`
	TranslatorID     int     `json:"translator_id"`
	IframeURL        string  `json:"iframe_url"`
	TrailerToken     string  `json:"trailer_token"`
	TrailerIframeURL string  `json:"trailer_iframe_url"`
	SeasonsCount     int     `json:"seasons_count"`
	EpisodesCount    int     `json:"episodes_count"`
	Category         string  `json:"category"`
	Age              int     `json:"age"`
	Countries        string  `json:"countries"`
	Actors           string  `json:"actors"`
	Directors        string  `json:"directors"`
	Studios          string  `json:"studios"`
	KinopoiskRating  float64 `json:"kinopoisk_rating"`
	KinopoiskVotes   int     `json:"kinopoisk_votes"`
	ImdbRating       float64 `json:"imdb_rating"`
	ImdbVotes        int     `json:"imdb_votes"`
}

// CreateRequest request params
type CreateRequest struct {
	Anime anime `json:"anime"`
}

// CreateResponse return struct in response
type CreateResponse struct {
	// API version
	Version string               `json:"v"`
	Anime   models.AnimeMoonWalk `json:"anime"`
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

	anime := models.AnimeMoonWalk{
		TitleRu:          r.Anime.TitleRu,
		TitleEn:          r.Anime.TitleEn,
		Year:             r.Anime.Year,
		Genres:           Strings.Split(r.Anime.Genres, ","),
		Poster:           r.Anime.Poster,
		Tagline:          r.Anime.Tagline,
		Description:      r.Anime.Description,
		Token:            r.Anime.Token,
		Type:             r.Anime.Type,
		KinopoiskID:      r.Anime.KinopoiskID,
		WorldArtID:       r.Anime.WorldArtID,
		Translator:       r.Anime.Translator,
		TranslatorID:     r.Anime.TranslatorID,
		IframeURL:        r.Anime.IframeURL,
		TrailerToken:     r.Anime.TrailerToken,
		TrailerIframeURL: r.Anime.TrailerIframeURL,
		SeasonsCount:     r.Anime.SeasonsCount,
		EpisodesCount:    r.Anime.EpisodesCount,
		Category:         r.Anime.Category,
		Age:              r.Anime.Age,
		Countries:        Strings.Split(r.Anime.Countries, ","),
		Actors:           Strings.Split(r.Anime.Actors, ","),
		Directors:        Strings.Split(r.Anime.Directors, ","),
		Studios:          Strings.Split(r.Anime.Studios, ","),
		KinopoiskRating:  r.Anime.KinopoiskRating,
		KinopoiskVotes:   r.Anime.KinopoiskVotes,
		ImdbRating:       r.Anime.ImdbRating,
		ImdbVotes:        r.Anime.ImdbVotes,
	}

	h.db.Create(&anime)

	c.JSON(http.StatusOK, CreateResponse{
		Version: "1",
		Anime:   anime,
	})
}
