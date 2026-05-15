package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateMilestone(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/milestones", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"iid":1,"title":"Sprint 1","state":"active"}`)
	})

	opts := &CreateMilestoneOptions{
		Title: Ptr("Sprint 1"),
	}
	ms, _, err := client.Milestones.CreateMilestone(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if ms.ID != 1 {
		t.Fatalf("expected ID=1, got %d", ms.ID)
	}
	if ms.Title != "Sprint 1" {
		t.Fatalf("expected title 'Sprint 1', got %q", ms.Title)
	}
	if ms.State != "active" {
		t.Fatalf("expected state 'active', got %q", ms.State)
	}
}

func TestEditMilestone(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/milestones/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"title":"Sprint 1 Updated"}`)
	})

	opts := &EditMilestoneOptions{
		Title: Ptr("Sprint 1 Updated"),
	}
	ms, _, err := client.Milestones.EditMilestone(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if ms.Title != "Sprint 1 Updated" {
		t.Fatalf("expected title 'Sprint 1 Updated', got %q", ms.Title)
	}
}

func TestListMilestones(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/milestones", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":1,"title":"Sprint 1"}]`)
	})

	milestones, resp, err := client.Milestones.ListMilestones(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(milestones) != 1 {
		t.Fatalf("expected 1 milestone, got %d", len(milestones))
	}
	if milestones[0].Title != "Sprint 1" {
		t.Fatalf("expected title 'Sprint 1', got %q", milestones[0].Title)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestGetMilestone(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/milestones/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"iid":1,"title":"Sprint 1","state":"active"}`)
	})

	ms, _, err := client.Milestones.GetMilestone(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if ms.ID != 1 {
		t.Fatalf("expected ID=1, got %d", ms.ID)
	}
	if ms.Title != "Sprint 1" {
		t.Fatalf("expected title 'Sprint 1', got %q", ms.Title)
	}
	if ms.State != "active" {
		t.Fatalf("expected state 'active', got %q", ms.State)
	}
}

func TestDeleteMilestone(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/milestones/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Milestones.DeleteMilestone(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListMilestoneIssues(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/milestones/1/issues", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"iid":1,"title":"Bug fix"}]`)
	})

	issues, _, err := client.Milestones.ListMilestoneIssues(context.Background(), 1, 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 1 {
		t.Fatalf("expected 1 issue, got %d", len(issues))
	}
	if issues[0].Title != "Bug fix" {
		t.Fatalf("expected title 'Bug fix', got %q", issues[0].Title)
	}
}
