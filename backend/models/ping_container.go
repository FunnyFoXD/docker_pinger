package models

import "time"

type ContainerPing struct {
	IP         uint      `json:"ip"`
	IPAddress  string    `json:"ip_address"`
	PingTime   int       `json:"ping_time"`
	LastUpdate time.Time `json:"last_update"`
}
