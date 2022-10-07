package time

import "time"

type Provider interface {
	Now() time.Time
}

type timeProvider struct {
}

func NewTimeProvider() Provider {
	return timeProvider{}
}

// Now returns the current time
func (t timeProvider) Now() time.Time {
	return time.Now()
}
