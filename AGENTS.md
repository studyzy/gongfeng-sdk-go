# AI Coding Guide for gongfeng-sdk-go

## Project Overview

Go SDK for Tencent Gongfeng (腾讯工蜂) REST API v3. Self-built HTTP client with zero heavy dependencies — only `github.com/google/go-querystring` for URL query encoding.

- **Module**: `github.com/studyzy/gongfeng-sdk-go`
- **Package**: `gongfeng` (single flat package, no sub-packages)
- **API Base**: `https://code.tencent.com/api/v3/`
- **Auth**: `PRIVATE-TOKEN` header
- **API Docs**: `docs/api/*.md` — reference these when implementing new endpoints

## Architecture

```
gongfeng.go          Core: Client, NewRequest, Do, Response, ErrorResponse
types.go             Shared models: User, Commit, Diff, Namespace, Member, Milestone, ListOptions, etc.
{resource}.go        One file per API resource: model + XxxService + methods
{resource}_test.go   Tests per resource
```

### Key types in gongfeng.go

- `Client` — holds baseURL, token, httpClient, and all Service fields
- `Response` — wraps `*http.Response` with pagination headers (X-Total, X-Page, etc.)
- `ErrorResponse` — implements `error`, returned for non-2xx responses
- `ClientOptionFunc` — functional options for `NewClient`

### Key types in types.go

- `ListOptions` — embed in List*Options for pagination (`Page`, `PerPage`)
- `User`, `Commit`, `Diff`, `Namespace`, `Member`, `Milestone` — shared across services
- `Ptr[T]` — generic helper to create pointer values for optional fields

## Code Conventions

### Service structure

Each `{resource}.go` file follows this layout:

```go
package gongfeng

import (
    "context"
    "fmt"
    "net/http"
)

// Model — match Gongfeng v3 JSON response fields exactly
type Branch struct {
    Name      string  `json:"name"`
    Protected bool    `json:"protected"`
    Commit    *Commit `json:"commit"`
}

// Service — single field: client *Client
type BranchesService struct {
    client *Client
}

// Options — pointer fields + omitempty for optional params
type CreateBranchOptions struct {
    BranchName *string `json:"branch_name,omitempty"`
    Ref        *string `json:"ref,omitempty"`
}

// Method — ctx first, three-value return
func (s *BranchesService) CreateBranch(ctx context.Context, pid interface{}, opts *CreateBranchOptions) (*Branch, *Response, error) {
    project, err := parseID(pid)
    if err != nil {
        return nil, nil, err
    }
    u := fmt.Sprintf("projects/%s/repository/branches", project)

    req, err := s.client.NewRequest(http.MethodPost, u, opts)
    if err != nil {
        return nil, nil, err
    }

    req = req.WithContext(ctx)

    var b Branch
    resp, err := s.client.Do(req, &b)
    if err != nil {
        return nil, resp, err
    }

    return &b, resp, nil
}
```

### Method signature rules

| Operation | Signature pattern |
| --- | --- |
| Get one | `GetXxx(ctx, pid, id) (*Xxx, *Response, error)` |
| List | `ListXxx(ctx, pid, opts) ([]*Xxx, *Response, error)` |
| Create | `CreateXxx(ctx, pid, opts) (*Xxx, *Response, error)` |
| Update | `UpdateXxx(ctx, pid, id, opts) (*Xxx, *Response, error)` |
| Delete | `DeleteXxx(ctx, pid, id) (*Response, error)` |

- **First param** is always `ctx context.Context`
- `pid interface{}` for project-scoped APIs (supports `int` or `"namespace/project"` string)
- Use `parseID(pid)` to convert to URL path segment
- Use `pathEscape(s)` for branch names, tag names, etc.
- Use `fmt.Sprintf` to build URL paths relative to `api/v3/`

### Options struct rules

- GET params: use `url:"field_name,omitempty"` struct tag (encoded by go-querystring)
- POST/PUT body: use `json:"field_name,omitempty"` struct tag (JSON encoded)
- All optional fields must be pointer types (`*string`, `*int`, `*bool`)
- Embed `ListOptions` for paginated list endpoints
- Naming: `{Action}{Resource}Options` (e.g., `CreateBranchOptions`, `ListProjectsOptions`)

### Request flow inside a method

```go
// 1. Parse resource ID
project, err := parseID(pid)

// 2. Build URL path (relative to /api/v3/)
u := fmt.Sprintf("projects/%s/repository/branches", project)

// 3. Create request (GET: opts → query string; POST/PUT: opts → JSON body)
req, err := s.client.NewRequest(http.MethodPost, u, opts)

// 4. Attach context
req = req.WithContext(ctx)

// 5. Execute and decode
var result Branch
resp, err := s.client.Do(req, &result)
```

### Error return convention

- On error building request: `return nil, nil, err`
- On error from Do: `return nil, resp, err` (always return resp so caller can inspect HTTP status)
- On success: `return &result, resp, nil`
- For delete operations: `return s.client.Do(req, nil)` (no response body)

### Naming conventions

- Service: `{Resource}sService` (e.g., `BranchesService`, `ProjectsService`)
- Receiver: `s` for all service methods (e.g., `func (s *BranchesService)`)
- Model: singular noun (e.g., `Branch`, `Project`, `MergeRequest`)
- Options: `{Action}{Resource}Options` (e.g., `CreateBranchOptions`)

### Comment conventions

- Every exported type: `// TypeName 描述。`
- Every exported func: `// FuncName 描述。`
- Package comment is in `gongfeng.go`: `// Package gongfeng ...`
- Use Chinese for descriptions (this is a Chinese-facing SDK)

## Adding a new API endpoint

1. Find the API spec in `docs/api/*.md`
2. If a new resource, create `{resource}.go`:
   - Define model struct(s) matching the JSON response
   - Define `{Resource}sService struct { client *Client }`
   - Add service field to `Client` in `gongfeng.go`
   - Initialize it in `NewClient()`
3. If extending existing resource, add to existing `{resource}.go`
4. Write methods following the patterns above
5. Add tests in `{resource}_test.go` using the `setup(t)` helper:

```go
func TestListBranches(t *testing.T) {
    client, mux := setup(t)

    mux.HandleFunc("/api/v3/projects/1/repository/branches", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            t.Fatalf("unexpected method: %s", r.Method)
        }
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, `[{"name":"main","protected":true}]`)
    })

    branches, _, err := client.Branches.ListBranches(context.Background(), 1, nil)
    if err != nil {
        t.Fatal(err)
    }
    if len(branches) != 1 || branches[0].Name != "main" {
        t.Fatalf("unexpected: %+v", branches)
    }
}
```

## Testing

```bash
go test ./...        # run all tests
go test -v ./...     # verbose
go test -run TestXxx # run specific test
```

- Use `setup(t)` from `gongfeng_test.go` to get a test `*Client` + `*http.ServeMux`
- Register handlers on mux with `/api/v3/...` paths
- Verify HTTP method, path, headers, and request body in handlers
- Test both success and error responses

## Build & Verify

```bash
go build ./...       # compile check
go vet ./...         # static analysis
go test ./...        # unit tests
goimports -w .       # format + fix imports
```

## Files you should NOT modify without understanding impact

- `gongfeng.go` — core HTTP machinery, changes affect all services
- `types.go` — shared models, changes affect multiple files
- `go.mod` — keep dependencies minimal
