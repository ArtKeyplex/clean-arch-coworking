package outbound

import (
	"context"

	"github.com/example/coworking/internal/booking/domain"
)

type BookingRepo interface {
	Save(ctx context.Context, b *domain.Booking) error
}
