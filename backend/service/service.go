// service/service.go

package service

import "context"

// BackendServiceServer is the server API for BackendService service.
type BackendServiceServer interface {
	// Ping receives a string and returns it back.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// PingRequest is the request message for the Ping method.
type PingRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

// PingResponse is the response message for the Ping method.
type PingResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}
