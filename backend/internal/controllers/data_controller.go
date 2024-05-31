package controllers

import (
	"encoding/json"
	"net/http"
)

// PingController handles ping requests.
type PingController struct {
	// grpcClient pb.MessageEchoServiceClient // Uncomment this line if you have a gRPC client
}

// PingHandler handles the ping endpoint.
func (c *PingController) PingHandler(w http.ResponseWriter, r *http.Request) {
    responseData := struct {
        Message string `json:"message"`
    }{
        Message: "pong",
    }

    jsonResponse, err := json.Marshal(responseData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
