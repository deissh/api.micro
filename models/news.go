package models

import "time"

type News struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Title      string `gorm:"not null" json:"title"`
	Annotation string `gorm:"not null" json:"annotation"`
	Body       string `gorm:"not null" json:"body"`
	Author     User   `default:"Anon" json:"author"`
	Preview    string `gorm:"not null" json:"preview"`
	Background string `default:"null" json:"background"`
	Types      string `default:"Системные" json:"types"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (n *News) View() News {
	// return news with private settings
	return News{
		Title:      n.Title,
		Annotation: n.Annotation,
		Author:     n.Author,
		Preview:    n.Preview,
		Background: n.Background,
		Types:      n.Types,
	}
}
