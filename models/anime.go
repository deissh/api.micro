package models

import (
	"errors"
	"github.com/lib/pq"
	"time"
)

// Translator contain id, name and episodes for this translator
type Translator struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	Token    string         `json:"moonwalk_token"`
	Episodes pq.StringArray `json:"episodes"`
}

// Anime main struct contain all props
type Anime struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	Title       string         `gorm:"not null;index:title" json:"title"`
	TitleEn     string         `gorm:"not null;index:title_en" json:"title_en"`
	TitleOr     string         `json:"title_or"`
	Annotation  string         `json:"annotation"`
	Description string         `json:"description"`
	Posters     pq.StringArray `gorm:"not null;type:varchar(2048)[]" json:"posters"`
	Type        string         `json:"type"`
	Genres      pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"genres"`
	Status      string         `json:"status"`
	Year        string         `json:"year"`
	Rating      float32        `gorm:"not null;default:5" json:"rating"`
	Votes       int            `gorm:"not null;default:0" json:"votes"`

	Translators []Translator `json:"translators"`

	WorldArtID  string `json:"world_art_id"`
	KinopoiskID string `json:"kinopoisk_id"`

	Countries pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"countries"`
	Actors    pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"actors"`
	Directors pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"directors"`
	Studios   pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"studios"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// AnimeShort contain only necessary information about this title
type AnimeShort struct {
	ID         uint           `gorm:"primary_key" json:"id"`
	Title      string         `gorm:"not null;index:title" json:"title"`
	TitleEn    string         `gorm:"not null;index:title_en" json:"title_en"`
	Annotation string         `json:"annotation"`
	Posters    pq.StringArray `gorm:"not null;type:varchar(2048)[]" json:"posters"`
	Type       string         `json:"type"`
	Genres     pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"genres"`
	Status     string         `json:"status"`
	Rating     float32        `json:"rating"`
	Votes      int            `json:"votes"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ViewShort return short information from full
func (a *Anime) ViewShort() AnimeShort {
	return AnimeShort{
		a.ID,
		a.Title,
		a.TitleEn,
		a.Annotation,
		a.Posters,
		a.Type,
		a.Genres,
		a.Status,
		a.Rating,
		a.Votes,
		a.CreatedAt,
		a.UpdatedAt,
	}
}

// GetEpisodesByTranslator return episodes by translator id
func (a *Anime) GetEpisodesByTranslator(id uint) ([]string, error) {
	for _, tr := range a.Translators {
		if id == tr.ID {
			return tr.Episodes, nil
		}
	}

	return nil, errors.New("not founded")
}

// AddVote add new vote to this anime
func (a *Anime) AddVote(value float32) {
	a.Votes++
	a.Rating = (a.Rating + value) / 2
}

// AddTranslator add new translators
func (a *Anime) AddTranslator(tr Translator) error {
	if len(tr.Episodes) == 0 {
		return errors.New("empty episodes")
	}
	a.Translators = append(a.Translators, tr)
	return nil
}
