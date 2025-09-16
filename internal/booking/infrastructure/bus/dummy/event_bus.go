package dummy

import (
	"context"

	"github.com/example/coworking/internal/booking/domain"
	"github.com/example/coworking/internal/booking/ports/outbound"
)

var _ outbound.EventBus = (*EventBus)(nil)

// EventBus is a non-functional in-memory event bus implementation that satisfies
// the outbound port for wiring the application service in tests or demos.
type EventBus struct{}

func (EventBus) Publish(ctx context.Context, events []domain.Event) error {
	return nil
}

// NewEventBus constructs a dummy event bus adapter.
func NewEventBus() EventBus {
	return EventBus{}
}
