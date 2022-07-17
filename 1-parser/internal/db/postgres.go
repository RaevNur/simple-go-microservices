package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DbHandler struct {
	DB *pgxpool.Pool
}

func InitPostgresDB(url string) (DbHandler, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return DbHandler{}, fmt.Errorf("cant parse config: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return DbHandler{}, fmt.Errorf("cant connect to db: %w", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		return DbHandler{}, fmt.Errorf("cant ping db: %w", err)
	}

	return DbHandler{db}, err
}
