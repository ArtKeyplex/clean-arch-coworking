package dummy

import (
	"context"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/domain"
	"github.com/example/coworking/internal/booking/ports/outbound"
)

var _ outbound.RoomSchedulePolicy = (*RoomSchedulePolicy)(nil)

// RoomSchedulePolicy is a no-op policy adapter that can be used to bootstrap
// the application without connecting to real availability or pricing systems.
type RoomSchedulePolicy struct{}

func (RoomSchedulePolicy) CheckAvailability(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) error {
	return nil
}

func (RoomSchedulePolicy) CalculatePrice(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) (domain.Money, error) {
	return domain.NewMoney(100, "USD"), nil
}

// NewPolicy constructs a dummy schedule policy adapter.
func NewPolicy() RoomSchedulePolicy {
	return RoomSchedulePolicy{}
}
