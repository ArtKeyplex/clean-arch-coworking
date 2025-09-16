package domain

import "errors"

var (
	ErrInvalidRange       = errors.New("invalid date range")
	ErrWrongState         = errors.New("booking in wrong state")
	ErrBookingNotFound    = errors.New("booking not found")
	ErrInvalidTransaction = errors.New("invalid payment transaction")
)
