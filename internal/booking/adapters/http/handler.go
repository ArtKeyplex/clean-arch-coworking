package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/ports/inbound"
)

type BookingHandler struct {
	svc inbound.BookingService
}

func NewBookingHandler(svc inbound.BookingService) *BookingHandler {
	return &BookingHandler{svc: svc}
}

func (h *BookingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RoomID string    `json:"room_id"`
		UserID string    `json:"user_id"`
		From   time.Time `json:"from"`
		To     time.Time `json:"to"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	roomID, _ := uuid.Parse(req.RoomID)
	userID, _ := uuid.Parse(req.UserID)
	id, err := h.svc.CreateBooking(r.Context(), roomID, userID, req.From, req.To)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}
