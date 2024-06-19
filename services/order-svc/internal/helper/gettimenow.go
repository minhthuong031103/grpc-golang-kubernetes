package helper

import "time"

// GetTimeNowInGMT7 returns the current time in GMT+7 formatted as a string
func GetTimeNowInGMT7() string {
	// Get the current local time
	now := time.Now()

	// Load the GMT+7 timezone
	location, err := time.LoadLocation("Asia/Bangkok") // GMT+7
	if err != nil {
		return "01-01-2024T00:00:00Z07:00"
	}

	// Convert the current time to GMT+7
	gmt7Time := now.In(location)

	// Format the time as a string
	timeString := gmt7Time.Format("2006-01-02T15:04:05Z07:00") // ISO 8601 format

	return timeString
}
