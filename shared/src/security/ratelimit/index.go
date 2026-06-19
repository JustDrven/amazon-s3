package ratelimit

import (
	"time"

	"justdrven.dev/storage/shared/src/pkg"
)

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		users: make(map[string]*UserStats),
	}
}

func (rl *RateLimiter) Allow(userID string) bool {
	rl.mu.Lock()
	stats, exists := rl.users[userID]
	if !exists {
		stats = &UserStats{}
		rl.users[userID] = stats
	}

	rl.mu.Unlock()
	stats.mu.Lock()
	defer stats.mu.Unlock()

	now := time.Now()
	if now.Before(stats.blockedUntil) {
		return false
	}

	cutoff := now.Add(-pkg.RATELIMIT_DURATION * time.Second)
	validRequests := stats.requestTimes[:0]
	for _, t := range stats.requestTimes {
		if t.After(cutoff) {
			validRequests = append(validRequests, t)
		}
	}

	stats.requestTimes = validRequests
	if len(stats.requestTimes) >= pkg.RATELIMIT_REQUESTS {
		stats.blockedUntil = now.Add(pkg.RATELIMIT_BLOCK_TIME * time.Minute)
		stats.requestTimes = nil
		return false
	}

	stats.requestTimes = append(stats.requestTimes, now)
	return true
}
