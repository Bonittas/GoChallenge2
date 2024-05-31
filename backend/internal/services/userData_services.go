package services

import (
	"context"
	"log"

	pb "github.com/Bonittas/GoChallenge2/proto" // Update the import path

	// "google.golang.org/grpc"
)

// PingService represents the gRPC server.
type PingService struct{}

// EchoMessage implements the EchoMessage gRPC method.
func (s *PingService) EchoMessage(ctx context.Context, req *pb.EchoMessageRequest) (*pb.EchoMessageResponse, error) {
	log.Printf("Received message: %s", req.Message)
	return &pb.EchoMessageResponse{Message: req.Message}, nil
}
