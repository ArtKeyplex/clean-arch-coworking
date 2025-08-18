package domain

import "github.com/google/uuid"

// BookingStatus defines the current state of booking
//go:generate stringer -type=BookingStatus

type BookingStatus int

const (
	Pending BookingStatus = iota
	Paid
	Cancelled
)

type Booking struct {
	id     uuid.UUID
	roomID uuid.UUID
	userID uuid.UUID
	slot   DateRange
	price  Money
	status BookingStatus
	events []Event
}

func NewBooking(roomID, userID uuid.UUID, slot DateRange, price Money) (*Booking, error) {
	if slot.IsZero() {
		return nil, ErrInvalidRange
	}
	b := &Booking{
		id:     uuid.New(),
		roomID: roomID,
		userID: userID,
		slot:   slot,
		price:  price,
		status: Pending,
	}
	b.raise(RoomBooked{BookingID: b.id.String(), RoomID: roomID.String(), UserID: userID.String()})
	return b, nil
}

func (b *Booking) ID() uuid.UUID { return b.id }

func (b *Booking) ConfirmPayment(txID string) error {
	if b.status != Pending {
		return ErrWrongState
	}
	b.status = Paid
	b.raise(BookingConfirmed{BookingID: b.id.String(), TxID: txID})
	return nil
}

func (b *Booking) PullEvents() []Event {
	ev := b.events
	b.events = nil
	return ev
}

func (b *Booking) raise(e Event) {
	b.events = append(b.events, e)
}
