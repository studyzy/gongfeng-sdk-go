package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCommit(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/commits/abc123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":"abc123","short_id":"abc1","title":"fix bug","message":"fix bug detail"}`)
	})

	commit, _, err := client.Commits.GetCommit(context.Background(), 1, "abc123")
	if err != nil {
		t.Fatal(err)
	}
	if commit.ID != "abc123" {
		t.Fatalf("expected commit ID 'abc123', got %q", commit.ID)
	}
	if commit.ShortID != "abc1" {
		t.Fatalf("expected short ID 'abc1', got %q", commit.ShortID)
	}
	if commit.Title != "fix bug" {
		t.Fatalf("expected title 'fix bug', got %q", commit.Title)
	}
}

func TestGetCommitDiff(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/commits/abc123/diff", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"old_path":"a.go","new_path":"a.go","diff":"@@ -1 +1 @@"}]`)
	})

	diffs, _, err := client.Commits.GetCommitDiff(context.Background(), 1, "abc123", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(diffs) != 1 {
		t.Fatalf("expected 1 diff, got %d", len(diffs))
	}
	if diffs[0].OldPath != "a.go" {
		t.Fatalf("expected old_path 'a.go', got %q", diffs[0].OldPath)
	}
}

func TestListCommitComments(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/commits/abc123/comments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"note":"looks good","path":"main.go","line":10}]`)
	})

	comments, _, err := client.Commits.ListCommitComments(context.Background(), 1, "abc123", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(comments) != 1 {
		t.Fatalf("expected 1 comment, got %d", len(comments))
	}
	if comments[0].Note != "looks good" {
		t.Fatalf("expected note 'looks good', got %q", comments[0].Note)
	}
	if comments[0].Line != 10 {
		t.Fatalf("expected line 10, got %d", comments[0].Line)
	}
}

func TestCreateCommitComment(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/commits/abc123/comments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"note":"looks good","path":"main.go","line":10}`)
	})

	comment, _, err := client.Commits.CreateCommitComment(context.Background(), 1, "abc123", &CreateCommitCommentOptions{
		Note: Ptr("looks good"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if comment.Note != "looks good" {
		t.Fatalf("expected note 'looks good', got %q", comment.Note)
	}
}

func TestListCommits(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/commits", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":"abc123","title":"fix bug"}]`)
	})

	commits, resp, err := client.Commits.ListCommits(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(commits) != 1 {
		t.Fatalf("expected 1 commit, got %d", len(commits))
	}
	if commits[0].ID != "abc123" {
		t.Fatalf("expected commit ID 'abc123', got %q", commits[0].ID)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestListCommitRefs(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/commits/abc123/refs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"type":"branch","name":"main"}]`)
	})

	refs, _, err := client.Commits.ListCommitRefs(context.Background(), 1, "abc123", nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(refs) != 1 {
		t.Fatalf("expected 1 ref, got %d", len(refs))
	}
	if refs[0].Type != "branch" {
		t.Fatalf("expected type 'branch', got %q", refs[0].Type)
	}
	if refs[0].Name != "main" {
		t.Fatalf("expected name 'main', got %q", refs[0].Name)
	}
}
