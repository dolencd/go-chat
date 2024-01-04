package common

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitializeConnection() (*pgx.Conn, error) {
	pg_url := os.Getenv("POSTGRES_URL")
	return pgx.Connect(context.Background(), pg_url)
}
