package monitor

import "time"

type MonitorStatus struct {
	Status    string        `json:"status"`
	Timestamp time.Duration `json:"timestamp"`
}
