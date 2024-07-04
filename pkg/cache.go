package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitRedis(addr, user, pw string, db int) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:        addr,
		Username:    user,
		Password:    pw,
		DB:          db,
		DialTimeout: time.Second * 5,
	})
	return c, c.Ping(context.Background()).Err()
}

type Cache struct {
	client *redis.Client
}

func NewCache(addr, user, pw string, db int) (*Cache, error) {
	client, err := InitRedis(addr, user, pw, db)
	if err != nil {
		return nil, err
	}
	return &Cache{client: client}, nil
}

func (c *Cache) CatchNil() error {
	if c == nil || c.client == nil {
		return fmt.Errorf("cache is nil")
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) string {
	if err := c.CatchNil(); err != nil {
		return ""
	}
	return c.client.Get(ctx, key).Val()

}

func (c *Cache) Set(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	if err := c.CatchNil(); err != nil {
		return err
	}
	return c.client.Set(ctx, key, val, timeout).Err()
}

func (c *Cache) IsExist(ctx context.Context, key string) bool {
	if err := c.CatchNil(); err != nil {
		return false
	}
	return c.client.Exists(ctx, key).Val() > 0
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	if err := c.CatchNil(); err != nil {
		return err
	}
	return c.client.Del(ctx, key).Err()
}
