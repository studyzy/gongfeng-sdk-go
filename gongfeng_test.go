package gongfeng

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// setup 创建一个测试用的 Client 和 ServeMux。
// 测试完成后自动关闭 server。
func setup(t *testing.T) (*Client, *http.ServeMux) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client, err := NewClient("test-token", WithBaseURL(server.URL))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return client, mux
}

func TestNewClientRequiresToken(t *testing.T) {
	_, err := NewClient("")
	if err == nil {
		t.Fatal("expected error when token is empty")
	}
}

func TestNewClientDefault(t *testing.T) {
	c, err := NewClient("my-token")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.baseURL.String() != "https://git.code.tencent.com/api/v3/" {
		t.Fatalf("unexpected baseURL: %s", c.baseURL.String())
	}
	if c.Projects == nil || c.Groups == nil || c.MergeRequests == nil {
		t.Fatal("services not initialized")
	}
}

func TestNewClientWithBaseURL(t *testing.T) {
	c, err := NewClient("my-token", WithBaseURL("https://git.example.com"))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.baseURL.String() != "https://git.example.com/api/v3/" {
		t.Fatalf("unexpected baseURL: %s", c.baseURL.String())
	}
}

func TestNewRequestSetsHeaders(t *testing.T) {
	c, _ := NewClient("secret-token")
	req, err := c.NewRequest(context.Background(), http.MethodGet, "projects", nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}

	if got := req.Header.Get("PRIVATE-TOKEN"); got != "secret-token" {
		t.Fatalf("expected PRIVATE-TOKEN=secret-token, got %q", got)
	}
	if got := req.Header.Get("User-Agent"); got != userAgent {
		t.Fatalf("expected User-Agent=%s, got %q", userAgent, got)
	}
}

func TestNewRequestGETWithOptions(t *testing.T) {
	c, _ := NewClient("token")
	opts := &ListOptions{Page: 2, PerPage: 50}
	req, err := c.NewRequest(context.Background(), http.MethodGet, "projects", opts)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}

	q := req.URL.Query()
	if q.Get("page") != "2" {
		t.Fatalf("expected page=2, got %q", q.Get("page"))
	}
	if q.Get("per_page") != "50" {
		t.Fatalf("expected per_page=50, got %q", q.Get("per_page"))
	}
}

func TestNewRequestPOSTWithBody(t *testing.T) {
	c, _ := NewClient("token")
	body := map[string]string{"name": "test-project"}
	req, err := c.NewRequest(context.Background(), http.MethodPost, "projects", body)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}

	if req.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("expected Content-Type=application/json, got %q", req.Header.Get("Content-Type"))
	}

	data, _ := io.ReadAll(req.Body)
	if !strings.Contains(string(data), "test-project") {
		t.Fatalf("body should contain 'test-project', got %q", string(data))
	}
}

func TestDoDecodesJSON(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "100")
		w.Header().Set("X-Page", "1")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	req, _ := client.NewRequest(context.Background(), http.MethodGet, "test", nil)
	var result map[string]string
	resp, err := client.Do(req, &result)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}

	if result["status"] != "ok" {
		t.Fatalf("unexpected result: %v", result)
	}
	if resp.TotalItems != 100 {
		t.Fatalf("expected TotalItems=100, got %d", resp.TotalItems)
	}
	if resp.CurrentPage != 1 {
		t.Fatalf("expected CurrentPage=1, got %d", resp.CurrentPage)
	}
}

func TestDoReturnsErrorResponse(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "404 Not Found"})
	})

	req, _ := client.NewRequest(context.Background(), http.MethodGet, "bad", nil)
	_, err := client.Do(req, nil)
	if err == nil {
		t.Fatal("expected error for 404")
	}

	errResp, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("expected *ErrorResponse, got %T", err)
	}
	if errResp.Response.StatusCode != 404 {
		t.Fatalf("expected status 404, got %d", errResp.Response.StatusCode)
	}
}

func TestDoWithContext(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	req, err := client.NewRequest(context.Background(), http.MethodGet, "health", nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}

	var result map[string]string
	_, err = client.Do(req, &result)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if result["status"] != "ok" {
		t.Fatalf("unexpected result: %v", result)
	}
}

func TestParseID(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
		wantErr  bool
	}{
		{42, "42", false},
		{"my-group/my-project", "my-group%2Fmy-project", false},
		{3.14, "", true},
	}

	for _, tt := range tests {
		got, err := parseID(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("parseID(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if got != tt.expected {
			t.Errorf("parseID(%v) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestPtr(t *testing.T) {
	s := Ptr("hello")
	if *s != "hello" {
		t.Fatalf("Ptr returned %q, want 'hello'", *s)
	}
	i := Ptr(42)
	if *i != 42 {
		t.Fatalf("Ptr returned %d, want 42", *i)
	}
}
