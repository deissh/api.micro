package models

import (
	"reflect"
	"testing"
	"time"
)

func TestNews_View(t *testing.T) {
	type fields struct {
		ID         uint
		Title      string
		Annotation string
		Body       string
		Author     User
		Preview    string
		Background string
		Types      string
		CreatedAt  time.Time
		UpdatedAt  time.Time
		DeletedAt  *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   News
	}{
		{
			name: "test suite 1",
			fields: fields{
				ID:         0,
				Title:      "test title",
				Annotation: "test ann",
				Body:       "",
				Author: User{
					ID:        1,
					FirstName: "Leon",
					LastName:  "Neff",
					Nickname:  "dev",
					Email:     "admin@google.com",
					Role:      "superadmin",
					Sex:       2,
				},
				Preview:    "https://some.url/cover.png",
				Background: "https://some.url/bg.png",
				Types:      "Системные",
			},
			want: News{
				ID:         0,
				Title:      "test title",
				Annotation: "test ann",
				Body:       "",
				Author: User{
					ID:        1,
					FirstName: "Leon",
					LastName:  "Neff",
					Nickname:  "dev",
					Email:     "admin@google.com",
					Role:      "superadmin",
					Sex:       2,
				},
				Preview:    "https://some.url/cover.png",
				Background: "https://some.url/bg.png",
				Types:      "Системные",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &News{
				ID:         tt.fields.ID,
				Title:      tt.fields.Title,
				Annotation: tt.fields.Annotation,
				Body:       tt.fields.Body,
				Author:     tt.fields.Author,
				Preview:    tt.fields.Preview,
				Background: tt.fields.Background,
				Types:      tt.fields.Types,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
				DeletedAt:  tt.fields.DeletedAt,
			}
			if got := n.View(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("News.View() = %v, want %v", got, tt.want)
			}
		})
	}
}
