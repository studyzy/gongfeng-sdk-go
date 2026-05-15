package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListTags(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/tags", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"name":"v1.0","message":"release"}]`)
	})

	tags, resp, err := client.Tags.ListTags(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) != 1 {
		t.Fatalf("expected 1 tag, got %d", len(tags))
	}
	if tags[0].Name != "v1.0" {
		t.Fatalf("expected tag name 'v1.0', got %q", tags[0].Name)
	}
	if tags[0].Message != "release" {
		t.Fatalf("expected tag message 'release', got %q", tags[0].Message)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestGetTag(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/tags/v1.0", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"v1.0","message":"release"}`)
	})

	tag, _, err := client.Tags.GetTag(context.Background(), 1, "v1.0")
	if err != nil {
		t.Fatal(err)
	}
	if tag.Name != "v1.0" {
		t.Fatalf("expected tag name 'v1.0', got %q", tag.Name)
	}
	if tag.Message != "release" {
		t.Fatalf("expected tag message 'release', got %q", tag.Message)
	}
}

func TestCreateTag(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/tags", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"v2.0","message":"new release"}`)
	})

	opts := &CreateTagOptions{
		TagName: Ptr("v2.0"),
		Ref:     Ptr("main"),
		Message: Ptr("new release"),
	}
	tag, _, err := client.Tags.CreateTag(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if tag.Name != "v2.0" {
		t.Fatalf("expected tag name 'v2.0', got %q", tag.Name)
	}
	if tag.Message != "new release" {
		t.Fatalf("expected tag message 'new release', got %q", tag.Message)
	}
}

func TestDeleteTag(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/tags/v1.0", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Tags.DeleteTag(context.Background(), 1, "v1.0")
	if err != nil {
		t.Fatal(err)
	}
}
