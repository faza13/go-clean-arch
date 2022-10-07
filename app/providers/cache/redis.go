package cache

import (
	"base/app/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"strconv"
	"time"
)

func NewRedisProvider(redisConfig *config.Cache) IProvider {
	address := fmt.Sprintf(
		"%s:%s",
		redisConfig.Host,
		redisConfig.Port,
	)

	defaultDB, err := strconv.Atoi(redisConfig.Database)

	if err != nil {
		defaultDB = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,              // address with port
		Password: redisConfig.Password, // no password set
		DB:       defaultDB,            // use default DB
	})

	return &Provider{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (p Provider) Set(key string, value interface{}, exp time.Duration) error {
	client := p.getClient()

	err := client.SetNX(p.ctx, key, value, exp)
	if err.Err() != nil {
		return err.Err()
	}

	return nil
}

func (p Provider) SetForever(key string, value interface{}) error {
	client := p.getClient()

	err := client.SetNX(p.ctx, key, value, redis.KeepTTL)

	if err.Err() != nil {
		return err.Err()
	}

	return nil
}

func (p Provider) Get(key string) (interface{}, error) {
	client := p.getClient()

	val, err := client.Get(p.ctx, key).Result()

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (p Provider) getClient() *redis.Client {
	redisClient := p.client.(redis.Client)

	return &redisClient
}
