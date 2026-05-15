package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListReleases(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/releases", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"tag_name":"v1.0","description":"First release"}]`)
	})

	releases, resp, err := client.Releases.ListReleases(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(releases) != 1 {
		t.Fatalf("expected 1 release, got %d", len(releases))
	}
	if releases[0].TagName != "v1.0" {
		t.Fatalf("expected tag_name 'v1.0', got %q", releases[0].TagName)
	}
	if releases[0].Description != "First release" {
		t.Fatalf("expected description 'First release', got %q", releases[0].Description)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestGetRelease(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/releases/v1.0", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"tag_name":"v1.0","description":"First release"}`)
	})

	release, _, err := client.Releases.GetRelease(context.Background(), 1, "v1.0")
	if err != nil {
		t.Fatal(err)
	}
	if release.TagName != "v1.0" {
		t.Fatalf("expected tag_name 'v1.0', got %q", release.TagName)
	}
	if release.Description != "First release" {
		t.Fatalf("expected description 'First release', got %q", release.Description)
	}
}

func TestCreateRelease(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/releases", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"tag_name":"v2.0","description":"New release"}`)
	})

	opts := &CreateReleaseOptions{
		TagName:     Ptr("v2.0"),
		Description: Ptr("New release"),
	}
	release, _, err := client.Releases.CreateRelease(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if release.TagName != "v2.0" {
		t.Fatalf("expected tag_name 'v2.0', got %q", release.TagName)
	}
	if release.Description != "New release" {
		t.Fatalf("expected description 'New release', got %q", release.Description)
	}
}

func TestUpdateRelease(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/releases/v1.0", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"tag_name":"v1.0","description":"Updated"}`)
	})

	opts := &UpdateReleaseOptions{
		Description: Ptr("Updated"),
	}
	release, _, err := client.Releases.UpdateRelease(context.Background(), 1, "v1.0", opts)
	if err != nil {
		t.Fatal(err)
	}
	if release.TagName != "v1.0" {
		t.Fatalf("expected tag_name 'v1.0', got %q", release.TagName)
	}
	if release.Description != "Updated" {
		t.Fatalf("expected description 'Updated', got %q", release.Description)
	}
}

func TestDeleteRelease(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/releases/v1.0", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Releases.DeleteRelease(context.Background(), 1, "v1.0")
	if err != nil {
		t.Fatal(err)
	}
}
