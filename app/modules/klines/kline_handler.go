package klines

import "github.com/jackc/pgx/v5"

type KlineHandler struct {
	service *KlineService
}

func NewKlineHandler(conn *pgx.Conn) *KlineHandler {
	service := NewKlineService(conn)
	return &KlineHandler{service: service}
}

func (kh *KlineHandler) GetKlines() (interface{}, error) {
	return kh.service.GetKlines()
}

func (kh *KlineHandler) InsertCandle(open float64, high float64, low float64, close float64, volume float64, time int) error {
	return kh.service.InsertCandle(open, high, low, close, volume, time)
}
