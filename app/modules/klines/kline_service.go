package klines

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quangvu30/go-socket/database"
	"github.com/quangvu30/go-socket/logger"
)

type KlineService struct {
	db *pgx.Conn
}

func NewKlineService(conn *pgx.Conn) *KlineService {
	return &KlineService{db: conn}
}

func (ks *KlineService) GetKlines() (pgx.Rows, error) {
	ctx := context.Background()
	logger.Log.Info("Getting klines from database")
	// Get klines from database
	res, err := ks.db.Query(ctx, "SELECT * FROM klines")
	if err != nil {
		logger.Log.Error("Error getting klines from database:", err)
		return nil, err
	}
	return res, nil
}

func (ks *KlineService) InsertCandle(open float64, high float64, low float64, close float64, volume float64, time int) error {
	ctx := context.Background()
	logger.Log.Info("Inserting candle into database")
	// Insert candle into database
	_, err := database.Conn.Exec(ctx, "INSERT INTO klines (open, high, low, close, volume, time) VALUES ($1, $2, $3, $4, $5, $6)", open, high, low, close, volume, time)
	if err != nil {
		logger.Log.Error("Error inserting candle into database:", err)
		return err
	}
	return nil
}
