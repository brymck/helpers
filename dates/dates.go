package dates

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	EndOfDayHour = 17
)

var (
	timezone *time.Location
)

type ProtoDate interface {
	GetYear() int32
	GetMonth() int32
	GetDay() int32
}

// Return a date in ISO format
func IsoFormat(t time.Time) string {
	return t.Format("2006-01-02")
}

func ProtoDateToTime(pd ProtoDate) time.Time {
	year := int(pd.GetYear())
	month := time.Month(pd.GetMonth())
	day := int(pd.GetDay())
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func LatestBusinessDate() time.Time {
	var err error
	timezone, err = time.LoadLocation("America/New_York")
	if err != nil {
		log.Warn("could not load America/New_York timezone, falling back to UTC")
		timezone = time.UTC
	}

	date := time.Now().In(timezone)
	switch date.Weekday() {
	case time.Monday:
		if date.Hour() < EndOfDayHour {
			date = date.AddDate(0, 0, -3)
		}
	case time.Sunday:
		date = date.AddDate(0, 0, -2)
	case time.Saturday:
		date = date.AddDate(0, 0, -1)
	default:
		if date.Hour() < EndOfDayHour {
			date = date.AddDate(0, 0, -1)
		}
	}
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
