package handlers

import (
	"github.com/jinzhu/gorm"
)

// ResponseData default error response
type ResponseData struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

// Handler contain all handlers and current db connection
type Handler struct {
	db *gorm.DB
}

// CreateHandlers create new handlers
func CreateHandlers(db *gorm.DB) Handler {
	return Handler{db}
}
