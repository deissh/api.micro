package main

import (
	"encoding/json"
	"github.com/deissh/api.micro/models"
	"github.com/deissh/api.micro/service-anime/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateRoute(t *testing.T) {
	router := SetupRouter()

	tests := []struct {
		name    string
		params  map[string]string
		status  int
		result  handlers.CreateResponse
		wantErr bool
	}{
		{
			"Create anime with empty params",
			map[string]string{},
			http.StatusBadRequest,
			handlers.CreateResponse{},
			true,
		},
		{
			"Create anime with invalid params",
			map[string]string{},
			http.StatusBadRequest,
			handlers.CreateResponse{
				Anime: models.Anime{
					Title:   "test",
					TitleEn: "",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := performRequest(router, "POST", "/anime.create", tt.params)

			if tt.wantErr {
				assert.Equal(t, tt.status, w.Code)
				return
			}

			assert.Equal(t, tt.status, w.Code)

			var response handlers.CreateResponse
			err := json.Unmarshal([]byte(w.Body.String()), &response)
			assert.Nil(t, err)

			assert.Equal(t, tt.result.Version, response.Version)
			assert.Equal(t, tt.result.Anime.Title, response.Anime.Title)
			assert.Equal(t, tt.result.Anime.TitleEn, response.Anime.TitleEn)
			assert.Equal(t, tt.result.Anime.TitleOr, response.Anime.TitleOr)
		})
	}
}
