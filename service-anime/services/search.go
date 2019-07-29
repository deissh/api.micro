package services

import (
	"github.com/nekko-ru/api/service-anime/models"
	"strings"
)

func (s Services) Search(query string, limit int, page int, sort string) ([]models.AnimeShort, error) {
	var animes []models.Anime
	var err error
	if len(query) > 0 {
		err = s.Db.Where(
			"to_tsvector(title) || to_tsvector(title_en) || to_tsvector(title_or) @@ plainto_tsquery(?)",
			query,
		).Limit(limit).Offset(limit * page).Order(strings.Replace(sort, "-", " desc", 1)).Find(&animes).Error
	} else {
		err = s.Db.Limit(limit).Offset(limit * page).Order(strings.Replace(sort, "-", " desc", 1)).Find(&animes).Error
	}

	if err != nil {
		s.Log.Error(err)
		return nil, err
	}

	var short_animes []models.AnimeShort
	for _, a := range animes {
		short_animes = append(short_animes, a.ViewShort())
	}

	return short_animes, nil
}
