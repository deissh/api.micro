package handlers

import (
	"github.com/nekko-ru/api/service-anime/services"
)

// ResponseData default error response
type ResponseData struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

// Handler contain all handlers and current db connection
type Handler struct {
	Srv services.Services
}
