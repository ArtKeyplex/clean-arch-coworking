package outbound

import (
	"context"
	"github.com/example/coworking/internal/booking/domain"
)

type EventBus interface {
	Publish(ctx context.Context, events []domain.Event) error
}
