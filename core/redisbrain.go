package core

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type redisBrain struct {
	pool *redis.Pool
}

func newRedisBrain() *redisBrain {
	return &redisBrain{
		pool: newPool(":6379"),
	}
}

func (b *redisBrain) remember(key string, value interface{}) error {
	c := b.pool.Get()
	defer c.Close()

	_, err := c.Do("SET", key, value)

	return err
}

func (b *redisBrain) recall(key string) (string, error) {
	c := b.pool.Get()
	defer c.Close()

	return redis.String(c.Do("GET", key))
}

func (b *redisBrain) rememberHash(hash string, key string, value interface{}) error {
	c := b.pool.Get()
	defer c.Close()

	_, err := c.Do("HSET", hash, key, value)

	return err
}

func (b *redisBrain) recallHash(hash string, key string) (string, error) {
	c := b.pool.Get()
	defer c.Close()

	return redis.String(c.Do("HGET", hash, key))
}

func (b *redisBrain) recallHashAll(hash string) ([]string, error) {
	c := b.pool.Get()
	defer c.Close()

	return redis.Strings(c.Do("HGETALL", hash))
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
