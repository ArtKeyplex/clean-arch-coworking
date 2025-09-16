package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/example/coworking/internal/booking/domain"
	"github.com/example/coworking/internal/booking/ports/inbound"
)

type BookingHandler struct {
	svc inbound.BookingService
}

func NewBookingHandler(svc inbound.BookingService) *BookingHandler {
	return &BookingHandler{svc: svc}
}

func (h *BookingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
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
	if req.RoomID == "" || req.UserID == "" {
		http.Error(w, "room_id and user_id are required", http.StatusBadRequest)
		return
	}
	roomID, err := uuid.Parse(req.RoomID)
	if err != nil {
		http.Error(w, "invalid room_id", http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	id, err := h.svc.CreateBooking(r.Context(), roomID, userID, req.From, req.To)
	if err != nil {
		writeError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"id": id.String()}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidRange), errors.Is(err, domain.ErrInvalidTransaction):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, domain.ErrWrongState):
		http.Error(w, err.Error(), http.StatusConflict)
	case errors.Is(err, domain.ErrBookingNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
