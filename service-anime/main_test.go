package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func performRequest(r http.Handler, method, path string, params map[string]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T) {
	router := SetupRouter()

	w := performRequest(router, "GET", "/_/ping", map[string]string{})
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
}
