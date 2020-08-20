package redis

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var (
	pool                      *redis.Pool
	errRedisAddrNotConfigured = errors.New("\"redis_addr\" configuration value not set")
)

// CreatePool initialises the Redis Connection Pool, not thread-safe
func CreatePool() error {
	addr := viper.GetString("redis_addr")
	if addr == "" {
		return errRedisAddrNotConfigured
	}
	pool = newPool(addr)
	return nil
}

// GetConnection acquires a Redis Connection from the Pool
// Thread Safe
func GetConnection() redis.Conn {
	return pool.Get()
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 5 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}
