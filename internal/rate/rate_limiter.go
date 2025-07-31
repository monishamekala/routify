package rate

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
)

type RateLimiter struct {
    rdb    *redis.Client
    limit  int           // max tokens
    window time.Duration // refill period
}

func NewRateLimiter(rdb *redis.Client, limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{rdb: rdb, limit: limit, window: window}
}

func (rl *RateLimiter) Allow(ctx context.Context, key string) (bool, error) {
    now := time.Now().Unix()
    pipe := rl.rdb.TxPipeline()

    tokensKey := "rate:" + key
    pipe.ZRemRangeByScore(ctx, tokensKey, "0", fmt.Sprint(now-int64(rl.window.Seconds())))
    pipe.ZCard(ctx, tokensKey)
    pipe.ZAdd(ctx, tokensKey, redis.Z{Score: float64(now), Member: now})
    pipe.Expire(ctx, tokensKey, rl.window)

    cmds, err := pipe.Exec(ctx)
    if err != nil {
        return false, err
    }

    used := cmds[1].(*redis.IntCmd).Val()
    return int(used) <= rl.limit, nil
}
