# Clean Arch Coworking

This repository contains a small skeleton implementing parts of a booking system for a coworking space. The code follows the ideas of Domain-Driven Design and clean (hexagonal) architecture.

The domain layer defines basic aggregates such as `Booking` and value objects like `DateRange` and `Money`. Application services expose use cases via ports and are wired with adapters in the `cmd` package.

Run the service:

```bash
 go run ./cmd/booking-service
```

Run unit tests:

```bash
 go test ./...
```
