// main.go

package main

import (
	"context"
	"log"
	"net"

	"github.com/Bonittas/GoChallenge2/backend/service"
	"google.golang.org/grpc"
)

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the service with the server
	service.RegisterBackendServiceServer(server, &backendServer{})

	// Listen for incoming connections
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Serve the gRPC server
	log.Println("Server listening on port 50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Define a struct to implement the BackendServiceServer interface generated by gRPC
type backendServer struct{}

// Implement the Ping method of the BackendServiceServer interface
func (s *backendServer) Ping(ctx context.Context, req *service.PingRequest) (*service.PingResponse, error) {
	return &service.PingResponse{Message: req.Message}, nil
}
