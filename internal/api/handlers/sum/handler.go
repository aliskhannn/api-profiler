package sum

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	Result int `json:"result"`
}

func (h *Handler) Sum(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	resp := Response{Result: a + b}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
