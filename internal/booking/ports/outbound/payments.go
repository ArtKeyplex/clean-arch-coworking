package outbound

import "context"

type PaymentGateway interface {
	Charge(ctx context.Context, bookingID string, amount int64, currency string) (string, error)
}
