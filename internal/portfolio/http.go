package portfolio

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service ServiceInterface
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/portfolio", h.GetAll)
	mux.HandleFunc("POST /api/portfolio", h.AddCoin)
	mux.HandleFunc("DELETE /api/portfolio", h.RemoveCoin)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	coins := h.service.GetAll()
	err := json.NewEncoder(w).Encode(coins)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) AddCoin(w http.ResponseWriter, r *http.Request) {
	var coin Coin
	err := json.NewDecoder(r.Body).Decode(&coin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.service.AddCoin(coin)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) RemoveCoin(w http.ResponseWriter, r *http.Request) {
	var id int32
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.service.RemoveCoin(id)
	w.WriteHeader(http.StatusNoContent)
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{service: service}
}
