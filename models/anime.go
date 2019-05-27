package models

import (
	"github.com/lib/pq"
	"time"
)

// ShortAnim return short info
type ShortAnim struct {
	ID uint `json:"id"`

	TitleRu     string         `json:"title_ru"`
	TitleEn     string         `json:"title_en"`
	Year        int            `json:"year"`
	Genres      pq.StringArray `json:"genres"`
	Poster      string         `json:"poster"`
	Tagline     string         `json:"tagline"`
	Description string         `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AnimeMoonWalk contain full data about anime
// fields: http://docs.moonwalk.cc/
type AnimeMoonWalk struct {
	// todo: set default values
	ID               uint           `gorm:"primary_key" json:"id"`
	TitleRu          string         `json:"title_ru"`
	TitleEn          string         `json:"title_en"`
	Year             int            `json:"year"`
	Genres           pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"genres"`
	Poster           string         `json:"poster"`
	Tagline          string         `json:"tagline"`
	Description      string         `json:"description"`
	Token            string         `json:"token"`
	Type             string         `json:"type"`
	KinopoiskID      int            `json:"kinopoisk_id"`
	WorldArtID       int            `json:"world_art_id"`
	Translator       string         `json:"translator"`
	TranslatorID     int            `json:"translator_id"`
	IframeURL        string         `json:"iframe_url"`
	TrailerToken     string         `json:"trailer_token"`
	TrailerIframeURL string         `json:"trailer_iframe_url"`
	SeasonsCount     int            `json:"seasons_count"`
	EpisodesCount    int            `json:"episodes_count"`
	Category         string         `json:"category"`
	Age              int            `json:"age"`
	Countries        pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"countries"`
	Actors           pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"actors"`
	Directors        pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"directors"`
	Studios          pq.StringArray `gorm:"not null;type:varchar(64)[]" json:"studios"`
	KinopoiskRating  float64        `json:"kinopoisk_rating"`
	KinopoiskVotes   int            `json:"kinopoisk_votes"`
	ImdbRating       float64        `json:"imdb_rating"`
	ImdbVotes        int            `json:"imdb_votes"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        *time.Time     `sql:"index" json:"-"`
}

type Episode struct {
	Name   string `json:"name"`
	Player string `json:"player"`
}

type Translator struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Episodes []Episode
}

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
	Rating      float32        `json:"rating"`
	Votes       int            `json:"votes"`

	EpisodesCount int          `gorm:"not null;default:0" json:"episodes_count"`
	Episodes      []Translator `json:"episodes"`

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

// ViewShort return view without some params
func (n *AnimeMoonWalk) ViewShort() ShortAnim {
	// return news with private settings
	return ShortAnim{
		n.ID,
		n.TitleRu,
		n.TitleEn,
		n.Year,
		n.Genres,
		n.Poster,
		n.Tagline,
		n.Description,
		n.CreatedAt,
		n.UpdatedAt,
	}
}
