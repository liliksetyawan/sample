package redis

import (
    "github.com/nexsoft-git/nexcommon/util"
    "github.com/go-redis/redis/v7"
)

func NewRedisUtil(
    nexsoft RedisConfig,
    gromart RedisConfig,
    etc RedisConfig,
) RedisUtil {
    result := &redisUtil{}

    result.nexsoft = util.ConnectRedis(
        util.NewRedisParam(
            nexsoft.Host,
            nexsoft.Port,
        ).
            DB(nexsoft.DB).
            MaxRetries(nexsoft.MaxRetries).
            Password(nexsoft.Password).
            Username(nexsoft.Username),
    )

    result.gromart = util.ConnectRedis(
        util.NewRedisParam(
            gromart.Host,
            gromart.Port,
        ).
            DB(gromart.DB).
            MaxRetries(gromart.MaxRetries).
            Password(gromart.Password).
            Username(gromart.Username),
    )

    result.etc = util.ConnectRedis(
        util.NewRedisParam(
            etc.Host,
            etc.Port,
        ).
            DB(etc.DB).
            MaxRetries(etc.MaxRetries).
            Password(etc.Password).
            Username(etc.Username),
    )

    return result
}

type redisUtil struct {
    nexsoft *redis.Client
    gromart *redis.Client
    etc *redis.Client
}

func (r *redisUtil) GetRedisNexsoftConnection() *redis.Client {
    return r.nexsoft
}

func (r *redisUtil) GetRedisGromartConnection() *redis.Client {
    return r.gromart
}

func (r *redisUtil) GetRedisEtcConnection() *redis.Client {
    return r.etc
}

func (r *redisUtil) Close() {
    _ = r.nexsoft.Close()
    _ = r.gromart.Close()
    _ = r.etc.Close()
}