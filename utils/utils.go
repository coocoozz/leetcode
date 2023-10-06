package utils

import "time"

func ElapsedTime(fn func()) time.Duration {
	now := time.Now()
	fn()
	return time.Now().Sub(now)
}
