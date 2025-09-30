package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/aliskhannn/api-profiler/internal/api/handlers/sum"
)

func main() {
	handler := sum.NewHandler()

	http.HandleFunc("/sum", handler.Sum)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
