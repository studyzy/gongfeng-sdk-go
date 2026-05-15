package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListUsers(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":1,"username":"admin","name":"Admin"}]`)
	})

	users, resp, err := client.Users.ListUsers(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}
	if users[0].Username != "admin" {
		t.Fatalf("expected username 'admin', got %q", users[0].Username)
	}
	if users[0].Name != "Admin" {
		t.Fatalf("expected name 'Admin', got %q", users[0].Name)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestGetUser(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/users/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"username":"admin","name":"Admin","email":"admin@test.com"}`)
	})

	user, _, err := client.Users.GetUser(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID != 1 {
		t.Fatalf("expected ID=1, got %d", user.ID)
	}
	if user.Username != "admin" {
		t.Fatalf("expected username 'admin', got %q", user.Username)
	}
	if user.Email != "admin@test.com" {
		t.Fatalf("expected email 'admin@test.com', got %q", user.Email)
	}
}

func TestGetCurrentUser(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"username":"me","name":"Me","email":"me@test.com"}`)
	})

	user, _, err := client.Users.GetCurrentUser(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if user.Username != "me" {
		t.Fatalf("expected username 'me', got %q", user.Username)
	}
	if user.Email != "me@test.com" {
		t.Fatalf("expected email 'me@test.com', got %q", user.Email)
	}
}
