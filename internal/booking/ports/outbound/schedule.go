package outbound

import (
	"context"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/domain"
)

type RoomSchedulePolicy interface {
	CheckAvailability(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) error
	CalculatePrice(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) (domain.Money, error)
}
