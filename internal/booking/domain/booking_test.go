package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewBooking(t *testing.T) {
	slot, _ := NewDateRange(time.Now(), time.Now().Add(time.Hour))
	b, err := NewBooking(uuid.New(), uuid.New(), slot, NewMoney(100, "USD"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(b.PullEvents()) == 0 {
		t.Errorf("expected event after creation")
	}
}
