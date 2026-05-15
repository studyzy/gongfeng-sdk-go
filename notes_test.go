package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateMergeRequestNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_requests/1/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"body":"LGTM"}`)
	})

	opts := &CreateMergeRequestNoteOptions{Body: Ptr("LGTM")}
	note, _, err := client.Notes.CreateMergeRequestNote(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if note.ID != 1 {
		t.Fatalf("expected ID=1, got %d", note.ID)
	}
	if note.Body != "LGTM" {
		t.Fatalf("expected body 'LGTM', got %q", note.Body)
	}
}

func TestListMergeRequestNotes(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_requests/1/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"body":"LGTM"}]`)
	})

	notes, _, err := client.Notes.ListMergeRequestNotes(context.Background(), 1, 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(notes) != 1 {
		t.Fatalf("expected 1 note, got %d", len(notes))
	}
	if notes[0].Body != "LGTM" {
		t.Fatalf("expected body 'LGTM', got %q", notes[0].Body)
	}
}

func TestGetMergeRequestNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_requests/1/notes/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"body":"LGTM"}`)
	})

	note, _, err := client.Notes.GetMergeRequestNote(context.Background(), 1, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if note.ID != 1 {
		t.Fatalf("expected ID=1, got %d", note.ID)
	}
	if note.Body != "LGTM" {
		t.Fatalf("expected body 'LGTM', got %q", note.Body)
	}
}

func TestUpdateMergeRequestNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_requests/1/notes/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"body":"Updated"}`)
	})

	opts := &UpdateMergeRequestNoteOptions{Body: Ptr("Updated")}
	note, _, err := client.Notes.UpdateMergeRequestNote(context.Background(), 1, 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if note.Body != "Updated" {
		t.Fatalf("expected body 'Updated', got %q", note.Body)
	}
}

func TestCreateIssueNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":2,"body":"Fixed"}`)
	})

	opts := &CreateIssueNoteOptions{Body: Ptr("Fixed")}
	note, _, err := client.Notes.CreateIssueNote(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if note.ID != 2 {
		t.Fatalf("expected ID=2, got %d", note.ID)
	}
	if note.Body != "Fixed" {
		t.Fatalf("expected body 'Fixed', got %q", note.Body)
	}
}

func TestListIssueNotes(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":2,"body":"Fixed"}]`)
	})

	notes, _, err := client.Notes.ListIssueNotes(context.Background(), 1, 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(notes) != 1 {
		t.Fatalf("expected 1 note, got %d", len(notes))
	}
	if notes[0].Body != "Fixed" {
		t.Fatalf("expected body 'Fixed', got %q", notes[0].Body)
	}
}

func TestGetIssueNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1/notes/2", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":2,"body":"Fixed"}`)
	})

	note, _, err := client.Notes.GetIssueNote(context.Background(), 1, 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if note.ID != 2 {
		t.Fatalf("expected ID=2, got %d", note.ID)
	}
}

func TestUpdateIssueNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/issues/1/notes/2", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":2,"body":"Updated note"}`)
	})

	opts := &UpdateIssueNoteOptions{Body: Ptr("Updated note")}
	note, _, err := client.Notes.UpdateIssueNote(context.Background(), 1, 1, 2, opts)
	if err != nil {
		t.Fatal(err)
	}
	if note.Body != "Updated note" {
		t.Fatalf("expected body 'Updated note', got %q", note.Body)
	}
}

func TestCreateReviewNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/reviews/1/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":3,"body":"review note"}`)
	})

	note, _, err := client.Notes.CreateReviewNote(context.Background(), 1, 1, &CreateReviewNoteOptions{Body: Ptr("review note")})
	if err != nil {
		t.Fatal(err)
	}
	if note.ID != 3 {
		t.Fatalf("expected ID=3, got %d", note.ID)
	}
}

func TestListReviewNotes(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/reviews/1/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":3,"body":"review note"}]`)
	})

	notes, _, err := client.Notes.ListReviewNotes(context.Background(), 1, 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(notes) != 1 {
		t.Fatalf("expected 1 note, got %d", len(notes))
	}
}

func TestGetReviewNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/reviews/1/notes/3", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":3,"body":"review note"}`)
	})

	note, _, err := client.Notes.GetReviewNote(context.Background(), 1, 1, 3)
	if err != nil {
		t.Fatal(err)
	}
	if note.ID != 3 {
		t.Fatalf("expected ID=3, got %d", note.ID)
	}
}

func TestUpdateReviewNote(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/reviews/1/notes/3", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":3,"body":"updated review note"}`)
	})

	note, _, err := client.Notes.UpdateReviewNote(context.Background(), 1, 1, 3, &UpdateReviewNoteOptions{Body: Ptr("updated review note")})
	if err != nil {
		t.Fatal(err)
	}
	if note.Body != "updated review note" {
		t.Fatalf("expected updated body, got %q", note.Body)
	}
}
