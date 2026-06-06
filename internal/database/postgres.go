package database

import (
	"context"
	"fmt"
	"log"

	"WebhookFlow/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(cfg *config.Config) *pgxpool.Pool {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	log.Println("database connected")

	return pool
}
