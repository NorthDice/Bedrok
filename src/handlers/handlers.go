package handlers

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	db     *sql.DB
	redis  *redis.Client
	logger *slog.Logger
}

func New(db *sql.DB, redis *redis.Client, logger *slog.Logger) *Handler {
	return &Handler{
		db:     db,
		redis:  redis,
		logger: logger,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Bedrok!")
}

func (h *Handler) Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Readiness(w http.ResponseWriter, r *http.Request) {
	if err := h.db.PingContext(r.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"status":"unavailable","db":"down"}`)
		return
	}

	if err := h.redis.Ping(r.Context()).Err(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"status":"unavailable","redis":"down"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok"}`)
}
