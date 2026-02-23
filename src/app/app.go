package app

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"bedrok/cnf"
	"bedrok/db"
	redisdb "bedrok/db/redis"
	"bedrok/handlers"
	"bedrok/logger"

	"github.com/redis/go-redis/v9"
)

type App struct {
	cfg    *cnf.Config
	db     *sql.DB
	redis  *redis.Client
	logger *slog.Logger
	router *http.ServeMux
}

func Init(ctx context.Context, cfg *cnf.Config) (*App, error) {
	log := logger.Init(cfg.Log)

	sqlDB, err := db.Init(ctx, cfg.DB)
	if err != nil {
		log.Error("Failed to connect to database", "error", err)
		return nil, err
	}

	redisClient, err := redisdb.Init(ctx, cfg.Redis)
	if err != nil {
		log.Error("Failed to connect to redis", "error", err)
		return nil, err
	}

	app := &App{
		cfg:    cfg,
		db:     sqlDB,
		redis:  redisClient,
		logger: log,
		router: http.NewServeMux(),
	}

	app.registerRoutes()

	return app, nil
}

func (a *App) registerRoutes() {
	h := handlers.New(a.db, a.redis, a.logger)
	a.router.HandleFunc("/", h.Home)
	a.router.HandleFunc("/healthz", h.Liveness)
	a.router.HandleFunc("/readyz", h.Readiness)
}

func (a *App) Router() http.Handler {
	return a.router
}

func (a *App) Close() {
	a.db.Close()
	a.redis.Close()
}
