package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/lib/pq"
)

func TestAnim_ViewShort(t *testing.T) {
	type fields struct {
		ID               uint
		TitleRu          string
		TitleEn          string
		Year             int
		Genres           pq.StringArray
		Poster           string
		Tagline          string
		Description      string
		Token            string
		Type             string
		KinopoiskID      int
		WorldArtID       int
		Translator       string
		TranslatorID     int
		IframeURL        string
		TrailerToken     string
		TrailerIframeURL string
		SeasonsCount     int
		EpisodesCount    int
		Category         string
		Age              int
		Countries        pq.StringArray
		Actors           pq.StringArray
		Directors        pq.StringArray
		Studios          pq.StringArray
		KinopoiskRating  float64
		KinopoiskVotes   int
		ImdbRating       float64
		ImdbVotes        int
		CreatedAt        time.Time
		UpdatedAt        time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   ShortAnim
	}{
		{
			name: "return short view",
			fields: fields{
				1,
				"test ru",
				"test en",
				2009,
				[]string{
					"1 genre",
					"2 genre",
				},
				"https://dev/cover.png",
				"some tagline",
				"Some desc",
				"external id",
				"sometype",
				141231,
				1234234,
				"some transleter",
				123,
				"https://dev/player.html",
				"external id",
				"https://dev/player.html",
				6,
				360,
				"SomeCategory",
				16,
				[]string{"Япония"},
				[]string{"Дзюнко Такэути", "Майли Флэнеген"},
				[]string{"Дзюнко Такэути", "Майли Флэнеген"},
				[]string{"Studio Pierrot Co. Ltd."},
				7.4,
				400,
				6.8,
				401,
				time.Unix(0, 0),
				time.Unix(1000, 0),
			},
			want: ShortAnim{
				1,
				"test ru",
				"test en",
				2009,
				[]string{
					"1 genre",
					"2 genre",
				},
				"https://dev/cover.png",
				"some tagline",
				"Some desc",
				time.Unix(0, 0),
				time.Unix(1000, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Anime{
				ID:               tt.fields.ID,
				TitleRu:          tt.fields.TitleRu,
				TitleEn:          tt.fields.TitleEn,
				Year:             tt.fields.Year,
				Genres:           tt.fields.Genres,
				Poster:           tt.fields.Poster,
				Tagline:          tt.fields.Tagline,
				Description:      tt.fields.Description,
				Token:            tt.fields.Token,
				Type:             tt.fields.Type,
				KinopoiskID:      tt.fields.KinopoiskID,
				WorldArtID:       tt.fields.WorldArtID,
				Translator:       tt.fields.Translator,
				TranslatorID:     tt.fields.TranslatorID,
				IframeURL:        tt.fields.IframeURL,
				TrailerToken:     tt.fields.TrailerToken,
				TrailerIframeURL: tt.fields.TrailerIframeURL,
				SeasonsCount:     tt.fields.SeasonsCount,
				EpisodesCount:    tt.fields.EpisodesCount,
				Category:         tt.fields.Category,
				Age:              tt.fields.Age,
				Countries:        tt.fields.Countries,
				Actors:           tt.fields.Actors,
				Directors:        tt.fields.Directors,
				Studios:          tt.fields.Studios,
				KinopoiskRating:  tt.fields.KinopoiskRating,
				KinopoiskVotes:   tt.fields.KinopoiskVotes,
				ImdbRating:       tt.fields.ImdbRating,
				ImdbVotes:        tt.fields.ImdbVotes,
				CreatedAt:        tt.fields.CreatedAt,
				UpdatedAt:        tt.fields.UpdatedAt,
			}
			if got := n.ViewShort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Anime.ViewShort() = %v, want %v", got, tt.want)
			}
		})
	}
}
