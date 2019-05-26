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
	// API version
	Version string `form:"v"`
	ID      string `form:"anime_id"`

	TitleRu          string   `json:"title_ru"`
	TitleEn          string   `json:"title_en"`
	Year             int      `json:"year"`
	Genres           []string `json:"genres"`
	Poster           string   `json:"poster"`
	Tagline          string   `json:"tagline"`
	Description      string   `json:"description"`
	Token            string   `json:"token"`
	Type             string   `json:"type"`
	KinopoiskID      int      `json:"kinopoisk_id"`
	WorldArtID       int      `json:"world_art_id"`
	Translator       string   `json:"translator"`
	TranslatorID     int      `json:"translator_id"`
	IframeURL        string   `json:"iframe_url"`
	TrailerToken     string   `json:"trailer_token"`
	TrailerIframeURL string   `json:"trailer_iframe_url"`
	SeasonsCount     int      `json:"seasons_count"`
	EpisodesCount    int      `json:"episodes_count"`
	Category         string   `json:"category"`
	Age              int      `json:"age"`
	Countries        []string `json:"countries"`
	Actors           []string `json:"actors"`
	Directors        []string `json:"directors"`
	Studios          []string `json:"studios"`
	KinopoiskRating  float64  `json:"kinopoisk_rating"`
	KinopoiskVotes   int      `json:"kinopoisk_votes"`
	ImdbRating       float64  `json:"imdb_rating"`
	ImdbVotes        int      `json:"imdb_votes"`

	AccessToken string `form:"access_token" binding:"required"`
}

// UpdateResponse response struct
type UpdateResponse struct {
	// API version
	Version string       `json:"v"`
	Anime   models.Anime `json:"anime"`
}

// UpdateAnime godoc
// @Summary update anime
// @Description Update anime and return it
// @ID update-anime
// @Accept  json
// @Produce  json
// @Param v query string false "service version"
// @Param anime_id query string true "anime id"
// @Param title query string false "title"
// @Param annotation query string false "annotation"
// @Param body query string false "body anime"
// @Param preview query string false "preview"
// @Param background query string false "background"
// @Param types query string false "anime types"
// @Param access_token query string true "user access_token"
// @Success 200 {object} handlers.UpdateRequest
// @Failure 400 {object} handlers.ResponseData
// @Router /anime.create [Get]
func (h Handler) UpdateAnime(c *gin.Context) {
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
		[]string{"animemaker", "admin", "superadmin"},
		[]string{"anime"},
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Unauthorized",
		})
		return
	}

	var anime models.Anime
	if err := h.db.First(&anime, r.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Anime did not find",
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
	if err := mergo.Merge(&r, anime); err != nil {
		c.JSON(http.StatusBadRequest, ResponseData{
			Status: http.StatusBadRequest,
			Data:   "Params error",
		})
		return
	}

	newAnime := models.Anime{
		//Types:      checkNull(anime.Types, r.Types),

		TitleRu:          r.TitleRu,
		TitleEn:          r.TitleEn,
		Year:             r.Year,
		Genres:           r.Genres,
		Poster:           r.Poster,
		Tagline:          r.Tagline,
		Description:      r.Description,
		Token:            r.Token,
		Type:             r.Type,
		KinopoiskID:      r.KinopoiskID,
		WorldArtID:       r.WorldArtID,
		Translator:       r.Translator,
		TranslatorID:     r.TranslatorID,
		IframeURL:        r.IframeURL,
		TrailerToken:     r.TrailerToken,
		TrailerIframeURL: r.TrailerIframeURL,
		SeasonsCount:     r.SeasonsCount,
		EpisodesCount:    r.EpisodesCount,
		Category:         r.Category,
		Age:              r.Age,
		Countries:        r.Countries,
		Actors:           r.Actors,
		Directors:        r.Directors,
		Studios:          r.Studios,
		KinopoiskRating:  r.KinopoiskRating,
		KinopoiskVotes:   r.KinopoiskVotes,
		ImdbRating:       r.ImdbRating,
		ImdbVotes:        r.ImdbVotes,
	}

	h.db.Create(&newAnime)

	c.JSON(http.StatusOK, UpdateResponse{
		Version: "1",
		Anime:   newAnime,
	})
}
