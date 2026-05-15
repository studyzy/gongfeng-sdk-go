package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateCommitStatus(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/commit/abc123/statuses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"sha":"abc123","status":"success","name":"ci"}`)
	})

	opts := &CreateCommitStatusOptions{
		State: Ptr("success"),
		Name:  Ptr("ci"),
	}
	status, _, err := client.CommitStatuses.CreateCommitStatus(context.Background(), 1, "abc123", opts)
	if err != nil {
		t.Fatal(err)
	}
	if status.ID != 1 {
		t.Fatalf("expected ID=1, got %d", status.ID)
	}
	if status.SHA != "abc123" {
		t.Fatalf("expected SHA 'abc123', got %q", status.SHA)
	}
	if status.Status != "success" {
		t.Fatalf("expected status 'success', got %q", status.Status)
	}
	if status.Name != "ci" {
		t.Fatalf("expected name 'ci', got %q", status.Name)
	}
}

func TestListCommitStatuses(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/commits/abc123/statuses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":1,"sha":"abc123","status":"success"}]`)
	})

	statuses, resp, err := client.CommitStatuses.ListCommitStatuses(context.Background(), 1, "abc123", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(statuses) != 1 {
		t.Fatalf("expected 1 status, got %d", len(statuses))
	}
	if statuses[0].SHA != "abc123" {
		t.Fatalf("expected SHA 'abc123', got %q", statuses[0].SHA)
	}
	if statuses[0].Status != "success" {
		t.Fatalf("expected status 'success', got %q", statuses[0].Status)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestGetCommitStatusResult(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/commits/main/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"sha":"abc123","status":"success","statuses":[{"id":1,"status":"success"}]}`)
	})

	result, _, err := client.CommitStatuses.GetCommitStatusResult(context.Background(), 1, "main")
	if err != nil {
		t.Fatal(err)
	}
	if result.SHA != "abc123" {
		t.Fatalf("expected SHA 'abc123', got %q", result.SHA)
	}
	if result.Status != "success" {
		t.Fatalf("expected status 'success', got %q", result.Status)
	}
	if len(result.Statuses) != 1 {
		t.Fatalf("expected 1 status entry, got %d", len(result.Statuses))
	}
}
