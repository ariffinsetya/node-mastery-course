package database

import (
	"log"

	goredis "github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type RedisDb struct {
	Client      *goredis.Client
	JsonHandler *rejson.Handler
}

func NewDatabase(address string, password string) (*RedisDb, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("goredis - failed to communicate to redis-server: %v", err)
		}
	}()
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client)
	return &RedisDb{
		Client:      client,
		JsonHandler: rh,
	}, nil
}
