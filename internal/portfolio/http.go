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
	coins, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(coins)
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
	err = h.service.AddCoin(coin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) RemoveCoin(w http.ResponseWriter, r *http.Request) {
	var id int32
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.RemoveCoin(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{service: service}
}
