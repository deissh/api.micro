package services

import (
	"github.com/nekko-ru/api/service-anime/models"
	"github.com/nekko-ru/api/service-anime/types"
)

func (s Services) Create(r types.CreateRequest) (models.Anime, error) {
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
		Translators: r.Translators,
		KinopoiskID: r.KinopoiskID,
		WorldArtID:  r.WorldArtID,
		Countries:   r.Countries,
		Actors:      r.Actors,
		Directors:   r.Directors,
		Studios:     r.Studios,
		Rating:      r.Rating,
		Votes:       r.Votes,
	}

	if err := s.Db.Create(&anime).Error; err != nil {
		s.Log.Debug(err)
		return models.Anime{}, err
	}
	return anime, nil
}
