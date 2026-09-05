package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(payload)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(
			w,
			http.StatusMethodNotAllowed,
			response{
				Status:  "error",
				Message: "method not allowed",
			},
		)

		return
	}

	writeJSON(
		w,
		http.StatusOK,
		response{
			Status:  "ok",
			Message: "service is running",
		},
	)
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger.Info(
			"request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}

func newRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	return mux
}

func main() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)

	handler := loggingMiddleware(
		logger,
		newRouter(),
	)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	logger.Info(
		"server starting",
		"address", server.Addr,
	)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(
			"server stopped unexpectedly",
			"error", err,
		)

		os.Exit(1)
	}
}