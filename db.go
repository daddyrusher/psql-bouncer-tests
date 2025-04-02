package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dsn := "postgres://user:password@localhost:5432/testdb?sslmode=disable"

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.TypeMap().RegisterDefaultPgType(&pgtype.UUID{}, "uuid")
		conn.TypeMap().RegisterDefaultPgType(&pgtype.Array[pgtype.UUID]{}, "uuid[]")
		return nil
	}

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
