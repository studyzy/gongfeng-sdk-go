package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListWatchers(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/watchers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":1,"username":"watcher1"}]`)
	})

	users, resp, err := client.Watchers.ListWatchers(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 watcher, got %d", len(users))
	}
	if users[0].Username != "watcher1" {
		t.Fatalf("expected username 'watcher1', got %q", users[0].Username)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestWatchProject(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/watchers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Watchers.WatchProject(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnwatchProject(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/watchers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Watchers.UnwatchProject(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
}
