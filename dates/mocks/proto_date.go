package mocks

import "time"

type MockProtoDate struct {
	time.Time
}

func (d MockProtoDate) GetYear() int32 {
	return int32(d.Year())
}

func (d MockProtoDate) GetMonth() int32 {
	return int32(d.Month())
}

func (d MockProtoDate) GetDay() int32 {
	return int32(d.Day())
}
