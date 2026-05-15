package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestAddWebhook(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/hooks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"url":"https://example.com/hook","push_events":true}`)
	})

	opts := &AddWebhookOptions{
		URL:        Ptr("https://example.com/hook"),
		PushEvents: Ptr(true),
	}
	hook, _, err := client.Webhooks.AddWebhook(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if hook.ID != 1 {
		t.Fatalf("expected ID=1, got %d", hook.ID)
	}
	if hook.URL != "https://example.com/hook" {
		t.Fatalf("expected URL 'https://example.com/hook', got %q", hook.URL)
	}
	if !hook.PushEvents {
		t.Fatal("expected PushEvents=true")
	}
}

func TestListWebhooks(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/hooks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":1,"url":"https://example.com/hook"}]`)
	})

	hooks, resp, err := client.Webhooks.ListWebhooks(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(hooks) != 1 {
		t.Fatalf("expected 1 hook, got %d", len(hooks))
	}
	if hooks[0].URL != "https://example.com/hook" {
		t.Fatalf("expected URL 'https://example.com/hook', got %q", hooks[0].URL)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestGetWebhook(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"url":"https://example.com/hook","push_events":true}`)
	})

	hook, _, err := client.Webhooks.GetWebhook(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if hook.ID != 1 {
		t.Fatalf("expected ID=1, got %d", hook.ID)
	}
	if hook.URL != "https://example.com/hook" {
		t.Fatalf("expected URL 'https://example.com/hook', got %q", hook.URL)
	}
	if !hook.PushEvents {
		t.Fatal("expected PushEvents=true")
	}
}

func TestEditWebhook(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"url":"https://example.com/hook2"}`)
	})

	opts := &EditWebhookOptions{
		URL: Ptr("https://example.com/hook2"),
	}
	hook, _, err := client.Webhooks.EditWebhook(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if hook.URL != "https://example.com/hook2" {
		t.Fatalf("expected URL 'https://example.com/hook2', got %q", hook.URL)
	}
}

func TestDeleteWebhook(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Webhooks.DeleteWebhook(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}
