package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateIssue(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"iid":1,"title":"Bug report","state":"opened"}`)
	})

	opts := &CreateIssueOptions{
		Title:       Ptr("Bug report"),
		Description: Ptr("Something is broken"),
	}
	issue, _, err := client.Issues.CreateIssue(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if issue.ID != 1 {
		t.Fatalf("expected ID=1, got %d", issue.ID)
	}
	if issue.Title != "Bug report" {
		t.Fatalf("expected title 'Bug report', got %q", issue.Title)
	}
	if issue.State != "opened" {
		t.Fatalf("expected state 'opened', got %q", issue.State)
	}
}

func TestUpdateIssue(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"iid":1,"title":"Updated bug","state":"opened"}`)
	})

	opts := &UpdateIssueOptions{
		Title: Ptr("Updated bug"),
	}
	issue, _, err := client.Issues.UpdateIssue(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if issue.Title != "Updated bug" {
		t.Fatalf("expected title 'Updated bug', got %q", issue.Title)
	}
}

func TestListIssues(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":1,"iid":1,"title":"Bug1"}]`)
	})

	issues, resp, err := client.Issues.ListIssues(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 1 {
		t.Fatalf("expected 1 issue, got %d", len(issues))
	}
	if issues[0].Title != "Bug1" {
		t.Fatalf("expected title 'Bug1', got %q", issues[0].Title)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestListUserIssues(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/issues", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"iid":1,"title":"Bug1"}]`)
	})

	issues, _, err := client.Issues.ListUserIssues(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 1 {
		t.Fatalf("expected 1 issue, got %d", len(issues))
	}
}

func TestGetIssue(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"iid":1,"title":"Bug1","state":"opened"}`)
	})

	issue, _, err := client.Issues.GetIssue(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if issue.ID != 1 {
		t.Fatalf("expected ID=1, got %d", issue.ID)
	}
	if issue.State != "opened" {
		t.Fatalf("expected state 'opened', got %q", issue.State)
	}
}

func TestDeleteIssue(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Issues.DeleteIssue(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetIssueSubscription(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1/subscribe", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `true`)
	})

	subscribed, _, err := client.Issues.GetIssueSubscription(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if !subscribed {
		t.Fatal("expected subscribed=true")
	}
}

func TestSubscribeIssue(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1/subscribe", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Issues.SubscribeIssue(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnsubscribeIssue(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Issues.UnsubscribeIssue(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}
