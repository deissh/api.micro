package services

import (
	"github.com/imdario/mergo"
	"github.com/nekko-ru/api/service-anime/models"
)

func (s Services) Update(id int, r models.Anime) (models.Anime, error) {
	var anime models.Anime
	if err := s.Db.Preload("Translators").First(&anime, id).Error; err != nil {
		s.Log.Error(err)
		return models.Anime{}, err
	}

	// merge two struct
	if err := mergo.Merge(&r, anime); err != nil {
		s.Log.Error(err)
		return models.Anime{}, err
	}

	s.Db.Where("a_id = ?", anime.ID).Unscoped().Delete(&r.Translators)
	s.Db.Model(&anime).Update(r).Save(&anime)

	return anime, nil
}
