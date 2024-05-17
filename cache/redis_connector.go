package cache

import (
	"context"

	"github.com/quangvu30/go-socket/app/models"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var ctx = context.Background()

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Set(key string, value string) {
	err := Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	val, err := Client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func ZAdd(key string, member redis.Z) {
	member.Member = models.Candle{
		Open:   member.Score,
		High:   member.Score,
		Low:    member.Score,
		Close:  member.Score,
		Volume: member.Score,
		Time:   int64(member.Score),
	}
	members := []redis.Z{member}
	err := Client.ZAdd(ctx, key, members...).Err()
	if err != nil {
		panic(err)
	}
}
