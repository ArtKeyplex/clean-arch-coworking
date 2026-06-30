package http

import (
	"context"
	"github.com/example/coworking/internal/booking/application"
	"github.com/example/coworking/internal/booking/domain"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockBookingService struct {
	application.BookingService // Встраиваем интерфейс
	cancelFunc                 func(ctx context.Context, id uuid.UUID) error
}

func (m *mockBookingService) CancelBooking(ctx context.Context, id uuid.UUID) error {
	return m.cancelFunc(ctx, id)
}

func TestBookingHandler_CancelBooking(t *testing.T) {
	id := uuid.New()

	tests := []struct {
		name           string
		bookingID      string
		mockErr        error
		expectedStatus int
	}{
		{
			name:           "success",
			bookingID:      id.String(),
			mockErr:        nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "error_not_found",
			bookingID:      id.String(),
			mockErr:        domain.ErrBookingNotFound,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "error_already_paid",
			bookingID:      id.String(),
			mockErr:        domain.ErrWrongState,
			expectedStatus: http.StatusConflict,
		},
		{
			name:           "invalid_uuid",
			bookingID:      "bad-id",
			mockErr:        nil,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			svc := &mockBookingService{
				cancelFunc: func(ctx context.Context, id uuid.UUID) error {
					return tt.mockErr
				},
			}
			h := &BookingHandler{svc: svc, logger: slog.Default()}

			req := httptest.NewRequest("POST", "/bookings/"+tt.bookingID+"/cancel", nil)
			req.SetPathValue("id", tt.bookingID)
			w := httptest.NewRecorder()

			h.CancelBooking(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
