package services

import (
	"github.com/nekko-ru/api/models"
)

func (service Services) Get(id int) (models.Anime, error) {
	var anime models.Anime
	if err := service.Db.Preload("Translators").First(&anime, id).Error; err != nil {
		return models.Anime{}, err
	}

	return anime, nil
}
