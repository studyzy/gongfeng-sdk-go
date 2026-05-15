package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListBranches(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/branches", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "2")
		fmt.Fprint(w, `[{"name":"main","protected":true},{"name":"dev","protected":false}]`)
	})

	branches, resp, err := client.Branches.ListBranches(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(branches) != 2 {
		t.Fatalf("expected 2 branches, got %d", len(branches))
	}
	if branches[0].Name != "main" {
		t.Fatalf("expected branch name 'main', got %q", branches[0].Name)
	}
	if !branches[0].Protected {
		t.Fatal("expected branch 'main' to be protected")
	}
	if resp.TotalItems != 2 {
		t.Fatalf("expected TotalItems=2, got %d", resp.TotalItems)
	}
}

func TestGetBranch(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/branches/main", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"main","protected":true}`)
	})

	branch, _, err := client.Branches.GetBranch(context.Background(), 1, "main")
	if err != nil {
		t.Fatal(err)
	}
	if branch.Name != "main" {
		t.Fatalf("expected branch name 'main', got %q", branch.Name)
	}
	if !branch.Protected {
		t.Fatal("expected branch to be protected")
	}
}

func TestCreateBranch(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/branches", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"feature","protected":false}`)
	})

	opts := &CreateBranchOptions{
		BranchName: Ptr("feature"),
		Ref:        Ptr("main"),
	}
	branch, _, err := client.Branches.CreateBranch(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if branch.Name != "feature" {
		t.Fatalf("expected branch name 'feature', got %q", branch.Name)
	}
	if branch.Protected {
		t.Fatal("expected branch not to be protected")
	}
}

func TestDeleteBranch(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/branches/feature", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Branches.DeleteBranch(context.Background(), 1, "feature")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetProtectedBranch(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/branches/main/protect", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"main","developers_can_push":false,"developers_can_merge":true}`)
	})

	pb, _, err := client.Branches.GetProtectedBranch(context.Background(), 1, "main")
	if err != nil {
		t.Fatal(err)
	}
	if pb.Name != "main" {
		t.Fatalf("expected name 'main', got %q", pb.Name)
	}
	if pb.DevelopersCanPush {
		t.Fatal("expected DevelopersCanPush=false")
	}
	if !pb.DevelopersCanMerge {
		t.Fatal("expected DevelopersCanMerge=true")
	}
}
