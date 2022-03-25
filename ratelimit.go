package main

import (
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
)

var userLastCalled = map[discord.UserID]time.Time{}

// checks to see if the author is rate limited
func isRateLimited(author discord.UserID) bool {
	lastCalled, exists := userLastCalled[author]
	return exists && time.Since(lastCalled) < rateLimitDuration
}

// updates the rate limit to the most recent access time
func updateRateLimit(author discord.UserID) {
	userLastCalled[author] = time.Now()
}
