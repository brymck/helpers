package dates

import "time"

// Return a date in ISO format
func IsoDate(t time.Time) string {
	return t.Format("2006-01-02")
}
