package main_test

import (
	"context"
	"testing"

	pb "github.com/Bonittas/GoChallenge2/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestGRPCPing(t *testing.T) {
	// Connect to the gRPC server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewPingServiceClient(conn)

	// Create a PingRequest
	req := &pb.PingRequest{Message: "test message"}

	// Call the Ping method on the server
	resp, err := client.Ping(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to call Ping RPC: %v", err)
	}

	// Assert the response
	assert.Equal(t, "test message", resp.Message)
}
