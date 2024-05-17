package main

import (
	"github.com/quangvu30/go-socket/app/modules/klines"
	"github.com/quangvu30/go-socket/config"
	"github.com/quangvu30/go-socket/database"
	"github.com/quangvu30/go-socket/logger"
)

func main() {
	config.GetConfig()
	println(config.ENV.PostgresUrl)
	database.Connect()
	logger.NewLogger()

	handler := klines.NewKlineHandler(database.Conn)
	handler.InsertCandle(0.01634790, 0.80000000, 0.01575800, 0.01577100, 148976.11427815, 1499644799999)
}
