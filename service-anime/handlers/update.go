package handlers

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/models"
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"net/http"
	Strings "strings"
)

// UpdateRequest request params
type UpdateRequest struct {
	Anime anime `json:"anime"`
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
// @Param anime body handlers.anime true "anime body"
// @Param anime_id query string true "anime_id"
// @Param v body query false "service version"
// @Param access_token query string true "user access_token"
// @Success 200 {object} handlers.UpdateRequest
// @Failure 400 {object} handlers.ResponseData
// @Router /anime.create [Post]
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
		c.DefaultQuery("access_token", ""),
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
	if err := h.db.First(&anime, c.DefaultQuery("anime_id", "")).Error; err != nil {
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
		//Types:      checkNull(anime.Types, r.Anime.Types),

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

	h.db.Create(&newAnime)

	c.JSON(http.StatusOK, UpdateResponse{
		Version: "1",
		Anime:   newAnime,
	})
}
