package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quangvu30/go-socket/config"
)

var Conn *pgx.Conn

func Connect() *pgx.Conn {
	var err error
	Conn, err = pgx.Connect(context.Background(), config.ENV.PostgresUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return Conn
}

func InitDatabase() {
	_, err := Conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS klines (
			id SERIAL PRIMARY KEY,
			open DOUBLE PRECISION,
			high DOUBLE PRECISION,
			low DOUBLE PRECISION,
			close DOUBLE PRECISION,
			volume DOUBLE PRECISION,
			time BIGINT
		);
	`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	}
}
