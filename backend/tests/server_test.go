package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bonittas/GoChallenge2/proto"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHTTPPing(t *testing.T) {
	router := mux.NewRouter()

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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}).Methods("POST", "OPTIONS")

	rr := httptest.NewRecorder()

	reqBody := []byte(`{"message": "test message"}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp proto.PingResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test message", resp.Message)
}

