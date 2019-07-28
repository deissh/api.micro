package services

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Services struct {
	Db  *gorm.DB
	Log *logrus.Logger
}
