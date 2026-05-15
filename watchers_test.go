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
		fmt.Fprint(w, `[{"project_id":1,"mute":false,"user":{"id":1,"username":"watcher1"}}]`)
	})

	watchers, resp, err := client.Watchers.ListWatchers(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(watchers) != 1 {
		t.Fatalf("expected 1 watcher, got %d", len(watchers))
	}
	if watchers[0].User == nil || watchers[0].User.Username != "watcher1" {
		t.Fatalf("expected username 'watcher1', got %+v", watchers[0].User)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestWatchProject(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/watch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"project_id":1,"mute":false,"user":{"id":1,"username":"watcher1"}}`)
	})

	watcher, _, err := client.Watchers.WatchProject(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if watcher.User == nil || watcher.User.Username != "watcher1" {
		t.Fatalf("unexpected watcher: %+v", watcher)
	}
}

func TestUnwatchProject(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/watch", func(w http.ResponseWriter, r *http.Request) {
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

func TestGetWatchStatus(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/watch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `true`)
	})

	watched, _, err := client.Watchers.GetWatchStatus(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if !watched {
		t.Fatal("expected watched=true")
	}
}
