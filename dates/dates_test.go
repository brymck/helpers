package dates

import (
	"testing"
	"time"

	"github.com/brymck/helpers/dates/mocks"
)

func TestIsoDate(t *testing.T) {
	date := time.Date(1984, 11, 23, 23, 59, 57, 0, time.UTC)
	actual := IsoDate(date)
	want := "1984-11-23"
	if actual != want {
		t.Errorf("want %q; got %q", want, actual)
	}
}

func TestProtoDateToTime(t *testing.T) {
	date := time.Date(1984, 11, 23, 23, 59, 57, 0, time.UTC)
	mockProtoDate := mocks.MockProtoDate{date}
	actual := ProtoDateToTime(mockProtoDate)
	want := time.Date(1984, 11, 23, 0, 0, 0, 0, time.UTC)
	if actual != want {
		t.Errorf("want %q; got %q", want, actual)
	}
}
