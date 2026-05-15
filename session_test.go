package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetSession(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"username":"user","email":"user@test.com","private_token":"secret"}`)
	})

	opts := &GetSessionOptions{
		Login:    Ptr("user"),
		Password: Ptr("pass"),
	}
	session, _, err := client.Session.GetSession(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if session.ID != 1 {
		t.Fatalf("expected ID=1, got %d", session.ID)
	}
	if session.Username != "user" {
		t.Fatalf("expected username 'user', got %q", session.Username)
	}
	if session.Email != "user@test.com" {
		t.Fatalf("expected email 'user@test.com', got %q", session.Email)
	}
	if session.PrivateToken != "secret" {
		t.Fatalf("expected private_token 'secret', got %q", session.PrivateToken)
	}
}
