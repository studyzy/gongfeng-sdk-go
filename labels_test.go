package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateLabel(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/labels", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"bug","color":"#FF0000"}`)
	})

	opts := &CreateLabelOptions{
		Name:  Ptr("bug"),
		Color: Ptr("#FF0000"),
	}
	label, _, err := client.Labels.CreateLabel(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if label.Name != "bug" {
		t.Fatalf("expected name 'bug', got %q", label.Name)
	}
	if label.Color != "#FF0000" {
		t.Fatalf("expected color '#FF0000', got %q", label.Color)
	}
}

func TestUpdateLabel(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/labels", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"bug","color":"#00FF00"}`)
	})

	opts := &UpdateLabelOptions{
		Name:  Ptr("bug"),
		Color: Ptr("#00FF00"),
	}
	label, _, err := client.Labels.UpdateLabel(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if label.Color != "#00FF00" {
		t.Fatalf("expected color '#00FF00', got %q", label.Color)
	}
}

func TestDeleteLabel(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/labels", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &DeleteLabelOptions{
		Name: Ptr("bug"),
	}
	_, err := client.Labels.DeleteLabel(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListLabels(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/labels", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"name":"bug","color":"#FF0000"}]`)
	})

	labels, resp, err := client.Labels.ListLabels(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(labels) != 1 {
		t.Fatalf("expected 1 label, got %d", len(labels))
	}
	if labels[0].Name != "bug" {
		t.Fatalf("expected name 'bug', got %q", labels[0].Name)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}
