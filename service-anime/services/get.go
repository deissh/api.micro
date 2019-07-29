package services

import (
	"github.com/nekko-ru/api/service-anime/models"
)

func (s Services) Get(id int) (models.Anime, error) {
	var anime models.Anime
	if err := s.Db.Preload("Translators").First(&anime, id).Error; err != nil {
		return models.Anime{}, err
	}

	return anime, nil
}
