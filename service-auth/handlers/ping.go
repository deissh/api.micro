package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ping struct {
	ID          string    `json:"id,omitempty"`
	ServiceName string    `json:"service,omitempty"`
	Time        time.Time `json:"time,omitempty"`
}

// PingResponse response struct
type PingResponse struct {
	// API version
	Version string `json:"v"`
	Ping    ping   `json:"ping"`
}

// PingCheck godoc
func (h Handler) PingCheck(c *gin.Context) {
	ping := ping{
		ID:          uuid.New().String(),
		ServiceName: "service-auth",
		Time:        time.Now().Local(),
	}

	c.JSON(http.StatusOK, PingResponse{
		Version: "1",
		Ping:    ping,
	})
}
