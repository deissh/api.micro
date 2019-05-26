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

// Anime contain full data about anime
// fields: http://docs.moonwalk.cc/
type Anime struct {
	// todo: set default values
	ID uint `gorm:"primary_key" json:"id"`

	TitleRu          string         `json:"title_ru"`
	TitleEn          string         `json:"title_en"`
	Year             int            `json:"year"`
	Genres           pq.StringArray `json:"genres"`
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
	Countries        pq.StringArray `json:"countries"`
	Actors           pq.StringArray `json:"actors"`
	Directors        pq.StringArray `json:"directors"`
	Studios          pq.StringArray `json:"studios"`
	KinopoiskRating  float64        `json:"kinopoisk_rating"`
	KinopoiskVotes   int            `json:"kinopoisk_votes"`
	ImdbRating       float64        `json:"imdb_rating"`
	ImdbVotes        int            `json:"imdb_votes"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// ViewShort return view without some params
func (n *Anime) ViewShort() ShortAnim {
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
