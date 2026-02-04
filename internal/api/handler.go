package api

import (
	"net/http"
	"strconv"
)

type Handler struct{}

func New() *Handler { return &Handler{} }

func (h *Handler) Sum(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	a, err := strconv.ParseInt(q.Get("a"), 10, 64)
	if err != nil {
		http.Error(w, "bad a", http.StatusBadRequest)
		return
	}
	b, err := strconv.ParseInt(q.Get("b"), 10, 64)
	if err != nil {
		http.Error(w, "bad b", http.StatusBadRequest)
		return
	}

	sum := a + b
	buf := make([]byte, 0, 32)
	buf = append(buf, `{"sum":`...)
	buf = strconv.AppendInt(buf, sum, 10)
	buf = append(buf, '}')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buf)
}

var jsonOK = []byte(`{"service":"api-profiler","ok":true}` + "\n")

func (h *Handler) JSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonOK)
}
