package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/aliskhannn/api-profiler/internal/api"
)

func main() {
	addr := env("ADDR", ":8080")

	mux := http.NewServeMux()
	h := api.New()
	mux.HandleFunc("GET /sum", h.Sum)
	mux.HandleFunc("GET /json", h.JSON)

	s := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	pprofAddr := env("PPROF_ADDR", "127.0.0.1:6060")

	go func() {
		log.Printf("pprof: http://%s/debug/pprof/\n", pprofAddr)
		if err := http.ListenAndServe(pprofAddr, nil); err != nil {
			log.Fatalf("pprof server: %v", err)
		}
	}()

	log.Printf("api: http://%s\n", addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("api server: %v", err)
	}
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
