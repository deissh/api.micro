package models

import (
	"github.com/lib/pq"
	"time"
)

type Token struct {
	ID           uint           `gorm:"primary_key" json:"id"`
	AccessToken  string         `gorm:"unique;not null" json:"access_token"`
	RefreshToken string         `gorm:"unique;not null" json:"refresh_token"`
	UserID       uint           `gorm:"not null" json:"user_id"`
	UserRole     string         `gorm:"not null;default:user" json:"role"`
	Permissions  pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"permissions"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
