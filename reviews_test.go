package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestInviteMRReviewer(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_request/1/reviewers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &InviteMRReviewerOptions{ReviewerID: Ptr(10)}
	_, err := client.Reviews.InviteMRReviewer(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveMRReviewer(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_request/1/reviewers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &RemoveMRReviewerOptions{ReviewerID: Ptr(10)}
	_, err := client.Reviews.RemoveMRReviewer(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApproveMR(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_request/1/review/approve", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &ApproveMROptions{Message: Ptr("LGTM")}
	_, err := client.Reviews.ApproveMR(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRejectMR(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_request/1/review/reject", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &RejectMROptions{Message: Ptr("needs work")}
	_, err := client.Reviews.RejectMR(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateCommitReview(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/commit_review", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"project_id":1,"title":"Review","state":"opened"}`)
	})

	opts := &CreateCommitReviewOptions{
		Title: Ptr("Review"),
	}
	review, _, err := client.Reviews.CreateCommitReview(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if review.ID != 1 {
		t.Fatalf("expected ID=1, got %d", review.ID)
	}
	if review.Title != "Review" {
		t.Fatalf("expected title 'Review', got %q", review.Title)
	}
	if review.State != "opened" {
		t.Fatalf("expected state 'opened', got %q", review.State)
	}
}

func TestListCommitReviews(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/commit_reviews", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"title":"Review"}]`)
	})

	reviews, _, err := client.Reviews.ListCommitReviews(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(reviews) != 1 {
		t.Fatalf("expected 1 review, got %d", len(reviews))
	}
	if reviews[0].Title != "Review" {
		t.Fatalf("expected title 'Review', got %q", reviews[0].Title)
	}
}

func TestGetCommitReview(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/commit_reviews/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"title":"Review","state":"opened"}`)
	})

	review, _, err := client.Reviews.GetCommitReview(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if review.ID != 1 {
		t.Fatalf("expected ID=1, got %d", review.ID)
	}
	if review.State != "opened" {
		t.Fatalf("expected state 'opened', got %q", review.State)
	}
}

func TestInviteCommitReviewer(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/commit_reviews/1/reviewers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &InviteCommitReviewerOptions{ReviewerID: Ptr(10)}
	_, err := client.Reviews.InviteCommitReviewer(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveCommitReviewer(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/commit_reviews/1/reviewers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &RemoveCommitReviewerOptions{ReviewerID: Ptr(10)}
	_, err := client.Reviews.RemoveCommitReviewer(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApproveCommitReview(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/commit_reviews/1/review/approve", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &ApproveCommitReviewOptions{Message: Ptr("approved")}
	_, err := client.Reviews.ApproveCommitReview(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRejectCommitReview(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/commit_reviews/1/review/reject", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &RejectCommitReviewOptions{Message: Ptr("rejected")}
	_, err := client.Reviews.RejectCommitReview(context.Background(), 1, 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}
