package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListNamespaces(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/namespaces", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":1,"name":"group1","path":"group1","kind":"group"}]`)
	})

	namespaces, resp, err := client.Namespaces.ListNamespaces(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(namespaces) != 1 {
		t.Fatalf("expected 1 namespace, got %d", len(namespaces))
	}
	if namespaces[0].Name != "group1" {
		t.Fatalf("expected name 'group1', got %q", namespaces[0].Name)
	}
	if namespaces[0].Path != "group1" {
		t.Fatalf("expected path 'group1', got %q", namespaces[0].Path)
	}
	if namespaces[0].Kind != "group" {
		t.Fatalf("expected kind 'group', got %q", namespaces[0].Kind)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}
