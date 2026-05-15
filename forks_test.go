package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestForkProject(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/fork/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":2,"name":"forked-project"}`)
	})

	project, _, err := client.Forks.ForkProject(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if project.ID != 2 {
		t.Fatalf("expected ID=2, got %d", project.ID)
	}
	if project.Name != "forked-project" {
		t.Fatalf("expected name 'forked-project', got %q", project.Name)
	}
}

func TestCreateForkRelation(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/fork/100", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Forks.CreateForkRelation(context.Background(), 1, 100)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteForkRelation(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/fork", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Forks.DeleteForkRelation(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
}
