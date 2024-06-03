// // service/service_test.go

// package service_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/Bonittas/GoChallenge2/backend/service"
// 	"github.com/stretchr/testify/assert"
// )

// func TestPing(t *testing.T) {
// 	svc := service.NewGRPCService()

// 	tests := []struct {
// 		name    string
// 		request *service.PingRequest
// 		want    string
// 	}{
// 		{
// 			name:    "Test with Hello",
// 			request: &service.PingRequest{Message: "Hello"},
// 			want:    "Hello",
// 		},
// 		{
// 			name:    "Test with World",
// 			request: &service.PingRequest{Message: "World"},
// 			want:    "World",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			resp, err := svc.Ping(context.Background(), tt.request)
// 			assert.NoError(t, err)
// 			assert.Equal(t, tt.want, resp.Message)
// 		})
// 	}
// }
