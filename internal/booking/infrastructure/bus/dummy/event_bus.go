package dummy

import (
	"context"

	"github.com/example/coworking/internal/booking/domain"
	"github.com/example/coworking/internal/booking/ports/outbound"
)

var _ outbound.EventBus = (*EventBus)(nil)

type EventBus struct{}

func (EventBus) Publish(ctx context.Context, events []domain.Event) error {
	return nil
}

func NewEventBus() EventBus {
	return EventBus{}
}
