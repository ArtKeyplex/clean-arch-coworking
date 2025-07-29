package application

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/domain"
	"github.com/example/coworking/internal/booking/ports/inbound"
	"github.com/example/coworking/internal/booking/ports/outbound"
)

var _ inbound.BookingService = (*Service)(nil)

type Service struct {
	repo   outbound.BookingRepo
	bus    outbound.EventBus
	policy outbound.RoomSchedulePolicy
}

func NewService(repo outbound.BookingRepo, bus outbound.EventBus, policy outbound.RoomSchedulePolicy) *Service {
	return &Service{repo: repo, bus: bus, policy: policy}
}

func (s *Service) CreateBooking(ctx context.Context, roomID, userID uuid.UUID, from, to time.Time) (uuid.UUID, error) {
	slot, err := domain.NewDateRange(from, to)
	if err != nil {
		return uuid.Nil, err
	}
	if err := s.policy.CheckAvailability(ctx, roomID, slot); err != nil {
		return uuid.Nil, err
	}
	price, err := s.policy.CalculatePrice(ctx, roomID, slot)
	if err != nil {
		return uuid.Nil, err
	}
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
	booking, err := s.repo.FindByID(ctx, bookingID)
	if err != nil {
		return err
	}
	if err := booking.ConfirmPayment(txID); err != nil {
		return err
	}
	if err := s.repo.Save(ctx, booking); err != nil {
		return err
	}
	return s.bus.Publish(ctx, booking.PullEvents())
}
