package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
}

type api struct {
	config config
	logger *slog.Logger
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "startup error: %v\n", err)
	}
}

func run() error {
	cfg := loadConfig()

	logger := newLogger(cfg.env)
	logger = logger.With("env", cfg.env, "component", "api")

	appCtx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM,
	)
	defer stop()

	api := &api{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      api.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	go func() {
		logger.Info("starting server", "addr", srv.Addr)
		if errSrv := srv.ListenAndServe(); errSrv != nil && errSrv != http.ErrServerClosed {
			logger.Error("HTTP server ListenAndServe error", "error", errSrv)
			stop()
		}
	}()

	<-appCtx.Done()
	logger.Info("shutdown initiated", "signal", "termination")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	var shutdownErr error

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("HTTP server shutdown error", "error", err)
		shutdownErr = err
	} else {
		logger.Info("HTTP server gracefully stopped.")
	}
	logger.Info("service exiting.")
	return shutdownErr
}

func loadConfig() config {
	var cfg config

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("invalid PORT env var: %v", err))
	}
	cfg.port = port

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	return cfg
}

func newLogger(env string) *slog.Logger {
	var handler slog.Handler
	if env == "development" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	return slog.New(handler)
}

type ApiStatus struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func (a *api) healthzHandler(w http.ResponseWriter, r *http.Request) {
	status := ApiStatus{
		Status:    "UP",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(status); err != nil {
		a.logger.Error("could not encode health status to JSON", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (a *api) routes() http.Handler {
	r := chi.NewRouter()

	// --- Middleware ---
	// Add some standard middleware for good practice.
	// Logger will log the start and end of each request with useful info.
	r.Use(middleware.Logger)
	// Recoverer will absorb panics and print the stack trace.
	r.Use(middleware.Recoverer)
	// RequestID sets a unique ID for each request.
	r.Use(middleware.RequestID)
	// RealIP gets the true client IP.
	r.Use(middleware.RealIP)

	r.Get("/healthz", a.healthzHandler)
	return r
}
