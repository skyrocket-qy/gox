package redis

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/skyrocket-qy/gox/lifecyclex"
)

func New(lc *lifecyclex.ConcurrentLifecycle, host, port, pass string) *redis.Client {
	log.Info().Msgf("start redis ")

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       0,
	})

	lc.Add(rdb, func(c context.Context) error {
		return rdb.Close()
	})

	return rdb
}

type DistributedLock struct {
	client *redis.Client
	key    string
	value  string
	ttl    time.Duration
}

func NewDistributedLock(client *redis.Client, key string, ttl time.Duration) *DistributedLock {
	return &DistributedLock{
		client: client,
		key:    key,
		ttl:    ttl,
	}
}

var NewUUID = uuid.NewString

// TryLock attempts to acquire the lock. Returns true if successful.
func (l *DistributedLock) TryLock(ctx context.Context) (bool, error) {
	l.value = NewUUID() // unique per holder
	ok, err := l.client.SetNX(ctx, l.key, l.value, l.ttl).Result()

	return ok, err
}

// Unlock releases the lock only if the value matches (safe unlock).
func (l *DistributedLock) Unlock(ctx context.Context) (bool, error) {
	// Lua script ensures atomic check-and-delete
	script := redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`)
	res, err := script.Run(ctx, l.client, []string{l.key}, l.value).Int()

	return res == 1, err
}
