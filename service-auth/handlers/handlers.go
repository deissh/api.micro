package handlers

import (
	"github.com/jinzhu/gorm"
)

// Error default response type
type ResponseData struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

// Handler contain all handlers and db connection
type Handler struct {
	db *gorm.DB
}

// CreateHandlers setup and return handlers with db connection
func CreateHandlers(db *gorm.DB) Handler {
	return Handler{db}
}
