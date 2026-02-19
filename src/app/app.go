package app

import (
	"bedrok/cnf"
	"bedrok/handlers"
	"bedrok/logger"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	cfg    *cnf.Config
	db     *sql.DB
	redis  *redis.Client
	logger *slog.Logger
	router *http.ServeMux
}

func New(cfg *cnf.Config) (*App, error) {
	log := logger.Init(cfg.Log)

	db, err := newDB(cfg.DB)
	if err != nil {
		log.Error("Failed to connect to database", "error", err)
		return nil, err
	}

	redisClient := newRedis(cfg.Redis)

	app := &App{
		cfg:    cfg,
		db:     db,
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
