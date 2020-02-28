package validation

import (
	"os"
	"testing"
)

func TestMustGetEnv(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		wantValue string
		wantPanic bool
	}{
		{"foo", "foo", "bar", "bar", false},
		{"foo", "foo", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv(tt.key, tt.value)
			if err != nil {
				t.Error(err)
			}
			if tt.wantPanic {
				assertPanic(t, func() {
					MustGetEnv(tt.key)
				})
			} else {
				value := MustGetEnv(tt.key)
				if value != tt.wantValue {
					t.Errorf("want %q; got %q", tt.wantValue, value)
				}
			}
		})
	}
}

func TestMustMapEnv(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     string
		wantValue string
		wantPanic bool
	}{
		{"foo", "foo", "bar", "bar", false},
		{"foo", "foo", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var value string
			err := os.Setenv(tt.key, tt.value)
			if err != nil {
				t.Error(err)
			}
			if tt.wantPanic {
				assertPanic(t, func() {
					MustMapEnv(&value, tt.key)
				})
			} else {
				MustMapEnv(&value, tt.key)
				if value != tt.wantValue {
					t.Errorf("want %q; got %q", tt.wantValue, value)
				}
			}
		})
	}
}
