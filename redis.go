package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"time"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/redis".
func init() {
	modules.Register("k6/x/redis", new(Redis))
}

// Redis is the k6 extension for a Redis client.
type Redis struct{}

// Client is the Redis client wrapper.
type Client struct {
	client *redis.Client
}

// XClient represents the Client constructor (i.e. `new redis.Client()`) and
// returns a new Redis client object.
func (r *Redis) XClient(ctxPtr *context.Context, server string, max int, timeout int) interface{} {
	opts := &redis.Options{Addr: server, PoolTimeout: time.Duration(timeout), PoolSize: max}
	rt := common.GetRuntime(*ctxPtr)
	return common.Bind(rt, &Client{client: redis.NewClient(opts)}, ctxPtr)
}

// Set the given key with the given value and expiration time.
func (c *Client) Set(key, value string, exp time.Duration) {
	c.client.Set(c.client.Context(), key, value, exp)
}

// Get returns the value for the given key.
func (c *Client) Get(key string) (string, error) {
	res, err := c.client.Get(c.client.Context(), key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
