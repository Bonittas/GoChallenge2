package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Bonittas/GoChallenge2/proto"
)

type server struct {
	proto.UnimplementedBackendServiceServer 
}

// Ping implements BackendServiceServer.Ping
func (s *server) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PingResponse, error) {
	log.Printf("Received message: %v", in.GetMessage())
	return &proto.PingResponse{Message: in.GetMessage()}, nil
}

func main() {
	go func() {
		listen, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		s := grpc.NewServer()
		proto.RegisterBackendServiceServer(s, &server{})
		reflection.Register(s)

		fmt.Println("gRPC server running on port :50051")
		if err := s.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		var req proto.PingRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := &proto.PingResponse{Message: req.Message}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("POST", "OPTIONS")

	handler := c.Handler(router)

	fmt.Println("HTTP server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
