package storage

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
    "github.com/redis/go-redis/v9"
    pb "github.com/yourusername/routify/proto"
)

var (
    ctx = context.Background()
    rdb *redis.Client
    db  *sql.DB
)

// InitClients initializes Redis and PostgreSQL clients
func InitClients(redisAddr, pgConnStr string) error {
    rdb = redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: "",
        DB:       0,
    })

    var err error
    db, err = sql.Open("postgres", pgConnStr)
    if err != nil {
        return err
    }

    return db.Ping()
}

// CacheKey generates a Redis key for a given origin-destination pair
func CacheKey(origin, destination string) string {
    return fmt.Sprintf("route:%s:%s", origin, destination)
}

// GetCachedRoute attempts to retrieve a route from Redis
func GetCachedRoute(origin, destination string) (*pb.RouteResponse, error) {
    val, err := rdb.Get(ctx, CacheKey(origin, destination)).Result()
    if err == redis.Nil {
        return nil, nil // Cache miss
    } else if err != nil {
        return nil, err
    }

    var resp pb.RouteResponse
    if err := json.Unmarshal([]byte(val), &resp); err != nil {
        return nil, err
    }

    return &resp, nil
}

// SaveRouteToCache stores a route in Redis with a TTL
func SaveRouteToCache(origin, destination string, resp *pb.RouteResponse) error {
    data, err := json.Marshal(resp)
    if err != nil {
        return err
    }

    return rdb.Set(ctx, CacheKey(origin, destination), data, 10*time.Minute).Err()
}

// SaveRouteToPostgres stores route history (simplified example)
func SaveRouteToPostgres(origin, destination string, totalTime float32) error {
    _, err := db.Exec(`
        INSERT INTO routes (origin, destination, total_time, created_at)
        VALUES ($1, $2, $3, NOW())`,
        origin, destination, totalTime)
    return err
}
