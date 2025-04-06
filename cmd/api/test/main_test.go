// filepath: /space/personal/learning/practice_go/cmd/api/main_test.go
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealthEndpoint(t *testing.T) {
	// Create a new router
	router := mux.NewRouter()
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is running"))
	})

	// Create a test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Send a GET request to the /api/health endpoint
	resp, err := http.Get(ts.URL + "/api/health")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	expectedBody := "API is running"
	if string(body) != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, string(body))
	}
}
