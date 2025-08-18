package domain

import "context"

// RoomSchedulePolicy checks for overlapping bookings and calculates price.
type RoomSchedulePolicy interface {
	Check(ctx context.Context, roomID string, slot DateRange) error
	CalculatePrice(roomID string, slot DateRange) Money
}
