package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/yourusername/routify/proto"
    "github.com/yourusername/routify/internal/routing"
)

const (
    port = ":50051"
)

grpcServer := grpc.NewServer(
    grpc.UnaryInterceptor(auth.UnaryJWTInterceptor()),
)


type server struct {
    pb.UnimplementedRoutePlannerServer
}

func (s *server) PlanRoute(ctx context.Context, req *pb.RouteRequest) (*pb.RouteResponse, error) {
    return routing.PlanRoute(req)
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterRoutePlannerServer(grpcServer, &server{})

    log.Printf("gRPC server listening on %s", port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
