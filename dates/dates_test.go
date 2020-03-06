package dates

import (
	"testing"
	"time"
)

func TestIsoDate(t *testing.T) {
	date := time.Date(1984, 11, 23, 23, 59, 57, 0, time.UTC)
	actual := IsoDate(date)
	want := "1984-11-23"
	if actual != want {
		t.Errorf("want %q; got %q", want, actual)
	}
}
