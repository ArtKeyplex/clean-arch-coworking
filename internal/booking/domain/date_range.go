package domain

import "time"

// DateRange represents a continuous time interval.
type DateRange struct {
	From time.Time
	To   time.Time
}

func NewDateRange(from, to time.Time) (DateRange, error) {
	if to.Before(from) || from.IsZero() || to.IsZero() {
		return DateRange{}, ErrInvalidRange
	}
	return DateRange{From: from, To: to}, nil
}

func (r DateRange) IsOverlapping(other DateRange) bool {
	return r.From.Before(other.To) && other.From.Before(r.To)
}

func (r DateRange) IsZero() bool {
	return r.From.IsZero() || r.To.IsZero()
}
