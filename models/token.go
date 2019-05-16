package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Token struct {
	gorm.Model
	AccessToken  string         `gorm:"unique;not null" json:"access_token"`
	RefreshToken string         `gorm:"unique;not null" json:"refresh_token"`
	UserId       int            `gorm:"not null" json:"user_id"`
	UserRole     string         `gorm:"not null;default:user" json:"role"`
	Permissions  pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"permissions"`
}
