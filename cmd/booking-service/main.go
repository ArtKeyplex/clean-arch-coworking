package main

import (
	"context"
	"log"
	"net/http"

	bookinghttp "github.com/example/coworking/internal/booking/adapters/http"
	"github.com/example/coworking/internal/booking/application"
	"github.com/example/coworking/internal/booking/domain"
)

type dummyRepo struct{}

func (d dummyRepo) Save(ctx context.Context, b *domain.Booking) error { return nil }

type dummyBus struct{}

func (d dummyBus) Publish(ctx context.Context, events []domain.Event) error { return nil }

type dummyPolicy struct{}

func (d dummyPolicy) Check(ctx context.Context, roomID string, slot domain.DateRange) error {
	return nil
}
func (d dummyPolicy) CalculatePrice(roomID string, slot domain.DateRange) domain.Money {
	return domain.NewMoney(100, "USD")
}

func main() {
	repo := dummyRepo{}
	bus := dummyBus{}
	policy := dummyPolicy{}
	svc := application.NewService(repo, bus, policy)
	handler := bookinghttp.NewBookingHandler(svc)

	log.Println("starting booking service on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
