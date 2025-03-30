package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"eduflow.dijam.io/internal/data"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Config holds the application configuration.
type Config struct {
	Port int    `env:"PORT" envDefault:"8080"`
	Env  string `env:"ENVIRONMENT" envDefault:"development"`
	DB   struct {
		DSN          string `env:"DB_DSN"`
		MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS" envDefault:"25"`
		MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS" envDefault:"25"`
		MaxIdleTime  string `env:"DB_MAX_IDLE_TIME" envDefault:"15m"`
	}
}

type application struct {
	logger *slog.Logger
	cfg    Config
	db     *sql.DB
	Models data.Models
}

func main() {
	// Load environment variables from the .env file.
	if err := godotenv.Load(); err != nil {
		slog.Info("Failed to find env file")
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		slog.Error("Failed to parse environment variables to config file")
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB(cfg)
	if err != nil {
		slog.Error(err.Error())
	}

	app := application{
		logger: logger,
		cfg:    cfg,
		db:     db,
		Models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.cfg.Port),
		Handler: app.routes(),
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 5 * time.Second,
		// IdleTimeout:  1 * time.Minute,
	}

	err = srv.ListenAndServe()

	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}

}

func openDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
