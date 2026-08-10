package cache

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis is a Cache backed by a Redis server.
//
// Every method swallows errors and degrades to a miss. A cache is an
// optimisation, and taking the application down because an optimisation is
// unreachable inverts the point of having one. Failures are logged at debug so
// a genuinely broken deployment is still diagnosable.
type Redis struct {
	client *redis.Client
	log    *slog.Logger
}

// Connect opens a client and verifies it. A bad URL is a configuration error
// and is returned; an unreachable server is not, because Redis may simply be
// starting alongside the API.
func Connect(ctx context.Context, url string, log *slog.Logger) (Cache, error) {
	if url == "" {
		return NoOp{}, nil
	}

	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	opts.DialTimeout = 3 * time.Second
	// Tight per-operation timeouts. A slow cache must never become slower than
	// recomputing the value it was meant to save.
	opts.ReadTimeout = 500 * time.Millisecond
	opts.WriteTimeout = 500 * time.Millisecond

	client := redis.NewClient(opts)

	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := client.Ping(pingCtx).Err(); err != nil {
		log.Warn("redis unreachable, continuing without cache", "err", err)
		_ = client.Close()
		return NoOp{}, nil
	}

	return &Redis{client: client, log: log}, nil
}

func (r *Redis) Enabled() bool { return true }

func (r *Redis) Get(ctx context.Context, key string) ([]byte, bool) {
	b, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Debug("cache get failed", "key", key, "err", err)
		}
		return nil, false
	}
	return b, true
}

func (r *Redis) Set(ctx context.Context, key string, value []byte, ttl time.Duration) {
	if err := r.client.Set(ctx, key, value, ttl).Err(); err != nil {
		r.log.Debug("cache set failed", "key", key, "err", err)
	}
}

// DeletePrefix removes every key under a prefix.
//
// SCAN rather than KEYS: KEYS blocks the server for the length of the scan, and
// a self-hosted deployment sharing Redis with anything else would feel it.
func (r *Redis) DeletePrefix(ctx context.Context, prefix string) {
	var cursor uint64
	for {
		keys, next, err := r.client.Scan(ctx, cursor, prefix+"*", 256).Result()
		if err != nil {
			r.log.Debug("cache scan failed", "prefix", prefix, "err", err)
			return
		}
		if len(keys) > 0 {
			if err := r.client.Del(ctx, keys...).Err(); err != nil {
				r.log.Debug("cache delete failed", "prefix", prefix, "err", err)
			}
		}
		if next == 0 {
			return
		}
		cursor = next
	}
}

func (r *Redis) Close() error { return r.client.Close() }
