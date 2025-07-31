package auth

import (
    "context"
    "errors"
    "strings"

    "github.com/golang-jwt/jwt/v5"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
)

var (
    JwtSecret = []byte("your-secret") // use env var or KMS in prod
)

func UnaryJWTInterceptor() grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {
        md, ok := metadata.FromIncomingContext(ctx)
        if !ok {
            return nil, errors.New("missing metadata")
        }

        tokens := md.Get("authorization")
        if len(tokens) == 0 {
            return nil, errors.New("authorization token required")
        }

        tokenStr := strings.TrimPrefix(tokens[0], "Bearer ")
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return JwtSecret, nil
        })
        if err != nil || !token.Valid {
            return nil, errors.New("invalid JWT token")
        }

        return handler(ctx, req)
    }
}
