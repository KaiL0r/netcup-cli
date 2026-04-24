package auth

import "time"

// ======================
// CLOCK ABSTRACTION
// ======================

type Clock interface {
	Now() time.Time
}

// production
type RealClock struct{}

func (RealClock) Now() time.Time {
	return time.Now()
}
