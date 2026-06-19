package ratelimit

import (
	"sync"
	"time"
)

type UserStats struct {
	mu           sync.Mutex
	requestTimes []time.Time
	blockedUntil time.Time
}

type RateLimiter struct {
	mu    sync.RWMutex
	users map[string]*UserStats
}
