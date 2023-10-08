package utils

import "time"

func ElapsedTime(fn func()) time.Duration {
	now := time.Now()
	fn()
	return time.Now().Sub(now)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
