package dates

import "time"

type ProtoDate interface {
	GetYear() int32
	GetMonth() int32
	GetDay() int32
}

// Return a date in ISO format
func IsoDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func ProtoDateToTime(pd ProtoDate) time.Time {
	year := int(pd.GetYear())
	month := time.Month(pd.GetMonth())
	day := int(pd.GetDay())
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
