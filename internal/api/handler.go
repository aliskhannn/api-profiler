package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct{}

func New() *Handler { return &Handler{} }

func (h *Handler) Sum(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	a, err := strconv.Atoi(q.Get("a"))
	if err != nil {
		http.Error(w, "bad a", http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(q.Get("b"))
	if err != nil {
		http.Error(w, "bad b", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(fmt.Sprintf(`{"sum":%d}`, a+b)))
}

func (h *Handler) JSON(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"service": "perf-http-api",
		"ok":      true,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
