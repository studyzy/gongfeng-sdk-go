package gongfeng

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestMRReviewOperations(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/merge_request/1/review/invite", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/api/v3/projects/1/merge_request/1/review/dismissals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/api/v3/projects/1/merge_request/1/review", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"id":1,"project_id":1,"state":"approving"}`)
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v3/projects/1/merge_request/1/reviewer/summary", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"project_id":1,"state":"approving"}`)
	})
	mux.HandleFunc("/api/v3/projects/1/merge_request/1/review/reopen", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"project_id":1,"state":"approving"}`)
	})
	mux.HandleFunc("/api/v3/projects/1/merge_request/1/review/cancel", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	if _, err := client.Reviews.InviteMRReviewer(context.Background(), 1, 1, &InviteMRReviewerOptions{ReviewerID: Ptr(10)}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Reviews.RemoveMRReviewer(context.Background(), 1, 1, &RemoveMRReviewerOptions{ReviewerID: Ptr(10)}); err != nil {
		t.Fatal(err)
	}
	review, _, err := client.Reviews.GetMRReview(context.Background(), 1, 1)
	if err != nil || review.ID != 1 {
		t.Fatalf("get mr review failed: %+v, %v", review, err)
	}
	review, _, err = client.Reviews.SubmitMRReviewSummary(context.Background(), 1, 1, &SubmitMRReviewSummaryOptions{ReviewerEvent: Ptr("approve"), Summary: Ptr("LGTM")})
	if err != nil || review.ID != 1 {
		t.Fatalf("submit mr review failed: %+v, %v", review, err)
	}
	review, _, err = client.Reviews.ReopenMRReview(context.Background(), 1, 1)
	if err != nil || review.ID != 1 {
		t.Fatalf("reopen mr review failed: %+v, %v", review, err)
	}
	if _, err := client.Reviews.CancelMRReview(context.Background(), 1, 1); err != nil {
		t.Fatal(err)
	}
}

func TestCommitReviewOperations(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/review", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"project_id":1,"title":"Review","state":"approving"}`)
	})
	mux.HandleFunc("/api/v3/projects/1/reviews", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"title":"Review"}]`)
	})
	mux.HandleFunc("/api/v3/projects/1/review/1", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet, http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"id":1,"title":"Review","state":"approving"}`)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v3/projects/1/review/1/invite", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/api/v3/projects/1/review/1/dismissals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/api/v3/projects/1/review/1/reviewer/summary", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"title":"Review","state":"approving"}`)
	})
	mux.HandleFunc("/api/v3/projects/1/review/1/reopen", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"title":"Review","state":"approving"}`)
	})
	mux.HandleFunc("/api/v3/projects/1/review/1/changed_files", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		fmt.Fprint(w, "review diff zip")
	})

	review, _, err := client.Reviews.CreateCommitReview(context.Background(), 1, &CreateCommitReviewOptions{Title: Ptr("Review")})
	if err != nil || review.ID != 1 {
		t.Fatalf("create commit review failed: %+v, %v", review, err)
	}
	reviews, _, err := client.Reviews.ListCommitReviews(context.Background(), 1, nil)
	if err != nil || len(reviews) != 1 {
		t.Fatalf("list commit reviews failed: %+v, %v", reviews, err)
	}
	review, _, err = client.Reviews.GetCommitReview(context.Background(), 1, 1)
	if err != nil || review.ID != 1 {
		t.Fatalf("get commit review failed: %+v, %v", review, err)
	}
	if _, err := client.Reviews.InviteCommitReviewer(context.Background(), 1, 1, &InviteCommitReviewerOptions{ReviewerID: Ptr(10)}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Reviews.RemoveCommitReviewer(context.Background(), 1, 1, &RemoveCommitReviewerOptions{ReviewerID: Ptr(10)}); err != nil {
		t.Fatal(err)
	}
	review, _, err = client.Reviews.SubmitCommitReviewSummary(context.Background(), 1, 1, &SubmitCommitReviewSummaryOptions{ReviewerEvent: Ptr("approve"), Summary: Ptr("ok")})
	if err != nil || review.ID != 1 {
		t.Fatalf("submit commit review failed: %+v, %v", review, err)
	}
	review, _, err = client.Reviews.ReopenCommitReview(context.Background(), 1, 1)
	if err != nil || review.ID != 1 {
		t.Fatalf("reopen commit review failed: %+v, %v", review, err)
	}
	review, _, err = client.Reviews.UpdateCommitReview(context.Background(), 1, 1, &UpdateCommitReviewOptions{Title: Ptr("Review")})
	if err != nil || review.ID != 1 {
		t.Fatalf("update commit review failed: %+v, %v", review, err)
	}

	var buf bytes.Buffer
	if _, err := client.Reviews.DownloadCommitReviewChangedFiles(context.Background(), 1, 1, &buf); err != nil {
		t.Fatal(err)
	}
	if buf.String() != "review diff zip" {
		t.Fatalf("unexpected body: %q", buf.String())
	}
}
