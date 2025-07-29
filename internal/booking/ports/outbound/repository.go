package outbound

import (
	"context"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/domain"
)

type BookingRepo interface {
	Save(ctx context.Context, b *domain.Booking) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Booking, error)
}
