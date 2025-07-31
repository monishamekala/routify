package main

import (
    "context"
    "log"
    "os"
    "time"

    "google.golang.org/grpc"
    pb "github.com/yourusername/routify/proto"
)

const (
    address = "localhost:50051"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewRoutePlannerClient(conn)

    origin := "StationA"
    destination := "StationB"
    if len(os.Args) == 3 {
        origin = os.Args[1]
        destination = os.Args[2]
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    resp, err := client.PlanRoute(ctx, &pb.RouteRequest{Origin: origin, Destination: destination})
    if err != nil {
        log.Fatalf("Route planning failed: %v", err)
    }

    log.Printf("Planned Route: %v", resp.Steps)
}
