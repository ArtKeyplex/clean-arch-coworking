package dummy

import (
	"context"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/domain"
	"github.com/example/coworking/internal/booking/ports/outbound"
)

var _ outbound.RoomSchedulePolicy = (*RoomSchedulePolicy)(nil)

type RoomSchedulePolicy struct{}

func (RoomSchedulePolicy) CheckAvailability(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) error {
	return nil
}

func (RoomSchedulePolicy) CalculatePrice(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) (domain.Money, error) {
	return domain.NewMoney(100, "USD"), nil
}

func NewPolicy() RoomSchedulePolicy {
	return RoomSchedulePolicy{}
}
