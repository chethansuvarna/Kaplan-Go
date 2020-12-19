package utils

import "time"

// GetTimeInFormat To get time in the given format
func GetTimeInFormat(format string, time time.Time) string {
	switch format {
	case "MM-DD-YYYY":
		return time.Format("01-02-2006")
	case "YYYY-MM-DD":
		return time.Format("2006-01-02")
	case "YYYY.MM.DD":
		return time.Format("2006.12.31")
	default:
		return time.Format("2006-01-02 15:04:05")
	}
}
