package dummy

import (
	"context"
	"github.com/example/coworking/internal/booking/application"
	"github.com/example/coworking/internal/booking/domain"
	"github.com/google/uuid"
)

// TODO: replace with real availability check against booking repository.
// Current implementation always returns nil, allowing double-bookings.
type AvailabilityChecker struct {
	repo application.BookingRepo
}

func NewAvailabilityChecker(repo application.BookingRepo) *AvailabilityChecker {
	return &AvailabilityChecker{repo: repo}
}
func (a *AvailabilityChecker) CheckAvailability(ctx context.Context, roomID uuid.UUID, slot domain.DateRange) error {
	bookings, err := a.repo.GetByRoom(ctx, roomID, slot)
	if err != nil {
		return err
	}
	for _, b := range bookings {
		if b.Status() == domain.Cancelled {
			continue

		}
		if b.Slot().IsOverlapping(slot) {
			return domain.ErrRoomNotAvailable
		}
	}
	return nil
}
