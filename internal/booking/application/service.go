package application

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/domain"
	"github.com/example/coworking/internal/booking/ports/inbound"
	"github.com/example/coworking/internal/booking/ports/outbound"
)

// Ensure Service implements inbound.BookingService
var _ inbound.BookingService = (*Service)(nil)

type Service struct {
	repo   outbound.BookingRepo
	bus    outbound.EventBus
	policy domain.RoomSchedulePolicy
}

func NewService(repo outbound.BookingRepo, bus outbound.EventBus, policy domain.RoomSchedulePolicy) *Service {
	return &Service{repo: repo, bus: bus, policy: policy}
}

func (s *Service) CreateBooking(ctx context.Context, roomID, userID uuid.UUID, from, to time.Time) (uuid.UUID, error) {
	slot, err := domain.NewDateRange(from, to)
	if err != nil {
		return uuid.Nil, err
	}
	if err := s.policy.Check(ctx, roomID.String(), slot); err != nil {
		return uuid.Nil, err
	}
	price := s.policy.CalculatePrice(roomID.String(), slot)
	booking, err := domain.NewBooking(roomID, userID, slot, price)
	if err != nil {
		return uuid.Nil, err
	}
	if err := s.repo.Save(ctx, booking); err != nil {
		return uuid.Nil, err
	}
	if err := s.bus.Publish(ctx, booking.PullEvents()); err != nil {
		return uuid.Nil, err
	}
	return booking.ID(), nil
}

func (s *Service) ConfirmPayment(ctx context.Context, bookingID uuid.UUID, txID string) error {
	// Repo should load booking; simplified as not implemented
	return nil
}
