package utils

import "time"

// Get a formatted Microseconds time
func GetMicTimeFormat() string {
	return time.Now().Format(KMicTimeFormat)
}
