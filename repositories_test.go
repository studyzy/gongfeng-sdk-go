package gongfeng

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestArchive(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/archive", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Write([]byte("archive-data"))
	})

	var buf bytes.Buffer
	_, err := client.Repositories.Archive(context.Background(), 1, &buf, nil)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() != "archive-data" {
		t.Fatalf("expected 'archive-data', got %q", buf.String())
	}
}

func TestListTree(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/tree", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":"abc","name":"main.go","type":"blob","mode":"100644"}]`)
	})

	nodes, _, err := client.Repositories.ListTree(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(nodes))
	}
	if nodes[0].Name != "main.go" {
		t.Fatalf("expected name 'main.go', got %q", nodes[0].Name)
	}
	if nodes[0].Type != "blob" {
		t.Fatalf("expected type 'blob', got %q", nodes[0].Type)
	}
}

func TestGetFile(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/files", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"file_name":"main.go","file_path":"main.go","size":100,"content":"base64data","ref":"main"}`)
	})

	opts := &GetFileOptions{
		FilePath: Ptr("main.go"),
		Ref:      Ptr("main"),
	}
	file, _, err := client.Repositories.GetFile(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if file.FileName != "main.go" {
		t.Fatalf("expected file_name 'main.go', got %q", file.FileName)
	}
	if file.Size != 100 {
		t.Fatalf("expected size=100, got %d", file.Size)
	}
	if file.Ref != "main" {
		t.Fatalf("expected ref 'main', got %q", file.Ref)
	}
}

func TestCreateFile(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/files", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"file_name":"new.go","file_path":"new.go"}`)
	})

	opts := &CreateFileOptions{
		FilePath:      Ptr("new.go"),
		BranchName:    Ptr("main"),
		Content:       Ptr("package main"),
		CommitMessage: Ptr("add new.go"),
	}
	file, _, err := client.Repositories.CreateFile(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if file.FileName != "new.go" {
		t.Fatalf("expected file_name 'new.go', got %q", file.FileName)
	}
}

func TestUpdateFile(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/files", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"file_name":"main.go","file_path":"main.go"}`)
	})

	opts := &UpdateFileOptions{
		FilePath:      Ptr("main.go"),
		BranchName:    Ptr("main"),
		Content:       Ptr("package main\n// updated"),
		CommitMessage: Ptr("update main.go"),
	}
	file, _, err := client.Repositories.UpdateFile(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if file.FileName != "main.go" {
		t.Fatalf("expected file_name 'main.go', got %q", file.FileName)
	}
}

func TestDeleteFile(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/files", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	opts := &DeleteFileOptions{
		FilePath:      Ptr("old.go"),
		BranchName:    Ptr("main"),
		CommitMessage: Ptr("remove old.go"),
	}
	_, err := client.Repositories.DeleteFile(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCompare(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/compare", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"commit":{"id":"abc"},"commits":[{"id":"abc"}],"diffs":[{"old_path":"a.go"}]}`)
	})

	opts := &CompareOptions{
		From: Ptr("main"),
		To:   Ptr("dev"),
	}
	result, _, err := client.Repositories.Compare(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if result.Commit == nil || result.Commit.ID != "abc" {
		t.Fatal("expected commit with id 'abc'")
	}
	if len(result.Commits) != 1 {
		t.Fatalf("expected 1 commit, got %d", len(result.Commits))
	}
	if len(result.Diffs) != 1 {
		t.Fatalf("expected 1 diff, got %d", len(result.Diffs))
	}
	if result.Diffs[0].OldPath != "a.go" {
		t.Fatalf("expected old_path 'a.go', got %q", result.Diffs[0].OldPath)
	}
}

func TestRepositoryRawContentAndCompareArchive(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/repository/blobs/abc123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		fmt.Fprint(w, "blob content")
	})
	mux.HandleFunc("/api/v3/projects/1/repository/commits/abc123/blob", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		fmt.Fprint(w, "commit blob content")
	})
	mux.HandleFunc("/api/v3/projects/1/repository/compare/changed_files", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		fmt.Fprint(w, "compare zip")
	})

	var blobBuf bytes.Buffer
	if _, err := client.Repositories.GetRawFile(context.Background(), 1, "abc123", &blobBuf, &GetRawFileOptions{FilePath: Ptr("main.go")}); err != nil {
		t.Fatal(err)
	}
	if blobBuf.String() != "blob content" {
		t.Fatalf("unexpected blob content: %q", blobBuf.String())
	}

	var commitBlobBuf bytes.Buffer
	if _, err := client.Repositories.GetCommitRawFile(context.Background(), 1, "abc123", &commitBlobBuf, &GetRawFileOptions{FilePath: Ptr("main.go")}); err != nil {
		t.Fatal(err)
	}
	if commitBlobBuf.String() != "commit blob content" {
		t.Fatalf("unexpected commit blob content: %q", commitBlobBuf.String())
	}

	var compareBuf bytes.Buffer
	if _, err := client.Repositories.DownloadCompareChangedFiles(context.Background(), 1, &compareBuf, &CompareOptions{From: Ptr("main"), To: Ptr("dev")}); err != nil {
		t.Fatal(err)
	}
	if compareBuf.String() != "compare zip" {
		t.Fatalf("unexpected compare content: %q", compareBuf.String())
	}
}
