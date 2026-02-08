package model

type PingResponse struct {
	Message     string `json:"message"`
	Uptime      int64  `json:"uptime"`       // in seconds
	CurrentTime string `json:"current_time"` // utc time string
}
