package gongfeng

import (
"context"
"encoding/json"
"net/http"
"net/http/httptest"
"testing"
)

func TestNewClientRequiresToken(t *testing.T) {
_, err := NewClient("", nil)
if err == nil {
t.Fatal("expected error when token is empty")
}
}

func TestCallUsesConfiguredBaseURLAndVersion(t *testing.T) {
server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
if r.URL.Path != "/api/v3/health" {
t.Fatalf("unexpected request path: %s", r.URL.Path)
}
if got := r.Header.Get("PRIVATE-TOKEN"); got != "test-token" {
t.Fatalf("unexpected token header: %q", got)
}

w.Header().Set("Content-Type", "application/json")
_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}))
defer server.Close()

	client, err := NewClient("test-token", &Options{BaseURL: server.URL, APIVersion: "v3"})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	var result map[string]string
	_, err = client.Call(context.Background(), http.MethodGet, "/health", nil, &result)
if err != nil {
t.Fatalf("Call returned error: %v", err)
}

if result["status"] != "ok" {
t.Fatalf("unexpected result: %#v", result)
}
}

func TestCallWithNilClient(t *testing.T) {
var cli *Client
_, err := cli.Call(context.Background(), http.MethodGet, "/health", nil, nil)
if err == nil {
t.Fatal("expected error for nil client")
}
}
