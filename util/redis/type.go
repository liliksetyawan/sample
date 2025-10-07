package redis

import (
    "github.com/go-redis/redis/v7"
)

type RedisUtil interface {
    GetRedisNexsoftConnection() *redis.Client
    GetRedisGromartConnection() *redis.Client
    GetRedisEtcConnection() *redis.Client
    Close()
}

type RedisConfig struct {
    Host       string
    Port       int
    DB         int
    MaxRetries int
    Password   string
    Username   string
}
