package helper

import "time"

// GetCurrentTime is a helper for return date  and hours now
func GetCurrentTime() string {
	time := time.Now().Local().Format("02_01_2006_15_04_05")
	return time
}
