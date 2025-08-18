package domain

import (
	"testing"
	"time"
)

func TestDateRangeOverlap(t *testing.T) {
	a, _ := NewDateRange(time.Now(), time.Now().Add(time.Hour))
	b, _ := NewDateRange(time.Now().Add(30*time.Minute), time.Now().Add(2*time.Hour))
	if !a.IsOverlapping(b) {
		t.Errorf("expected overlap")
	}
}
