package outbound

import (
	"context"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/domain"
)

// RoomSchedulePolicy provides access to availability checks and pricing rules
// for coworking rooms. Infrastructure adapters implement this interface and
// the application layer depends on it via the outbound port to keep the domain
// isolated from persistence concerns.
type RoomSchedulePolicy interface {
	CheckAvailability(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) error
	CalculatePrice(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) (domain.Money, error)
}
