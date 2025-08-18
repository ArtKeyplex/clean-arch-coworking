package inbound

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type BookingService interface {
	CreateBooking(ctx context.Context, roomID, userID uuid.UUID, from, to time.Time) (uuid.UUID, error)
	ConfirmPayment(ctx context.Context, bookingID uuid.UUID, txID string) error
}
