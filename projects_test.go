package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListProjects(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "2")
		fmt.Fprint(w, `[{"id":1,"name":"proj1"},{"id":2,"name":"proj2"}]`)
	})

	projects, resp, err := client.Projects.ListProjects(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) != 2 {
		t.Fatalf("expected 2 projects, got %d", len(projects))
	}
	if projects[0].ID != 1 {
		t.Fatalf("expected project ID 1, got %d", projects[0].ID)
	}
	if projects[0].Name != "proj1" {
		t.Fatalf("expected project name 'proj1', got %q", projects[0].Name)
	}
	if resp.TotalItems != 2 {
		t.Fatalf("expected TotalItems=2, got %d", resp.TotalItems)
	}
}

func TestGetProject(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"name":"my-project","path":"my-project","default_branch":"main"}`)
	})

	project, _, err := client.Projects.GetProject(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if project.ID != 1 {
		t.Fatalf("expected project ID 1, got %d", project.ID)
	}
	if project.Name != "my-project" {
		t.Fatalf("expected name 'my-project', got %q", project.Name)
	}
	if project.DefaultBranch != "main" {
		t.Fatalf("expected default_branch 'main', got %q", project.DefaultBranch)
	}
}

func TestCreateProject(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":3,"name":"new-proj"}`)
	})

	opts := &CreateProjectOptions{
		Name: Ptr("new-proj"),
	}
	project, _, err := client.Projects.CreateProject(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if project.ID != 3 {
		t.Fatalf("expected project ID 3, got %d", project.ID)
	}
	if project.Name != "new-proj" {
		t.Fatalf("expected name 'new-proj', got %q", project.Name)
	}
}

func TestSearchProjects(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/search/test", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"name":"test-proj"}]`)
	})

	projects, _, err := client.Projects.SearchProjects(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(projects))
	}
	if projects[0].Name != "test-proj" {
		t.Fatalf("expected name 'test-proj', got %q", projects[0].Name)
	}
}

func TestListProjectMembers(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":10,"username":"user1","access_level":30}]`)
	})

	members, _, err := client.Projects.ListProjectMembers(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 1 {
		t.Fatalf("expected 1 member, got %d", len(members))
	}
	if members[0].ID != 10 {
		t.Fatalf("expected member ID 10, got %d", members[0].ID)
	}
	if members[0].AccessLevel != DeveloperPermission {
		t.Fatalf("expected access level %d, got %d", DeveloperPermission, members[0].AccessLevel)
	}
}

func TestAddProjectMember(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":10,"username":"user1","access_level":30}`)
	})

	opts := &AddProjectMemberOptions{
		UserID:      Ptr(10),
		AccessLevel: Ptr(DeveloperPermission),
	}
	member, _, err := client.Projects.AddProjectMember(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if member.ID != 10 {
		t.Fatalf("expected member ID 10, got %d", member.ID)
	}
	if member.Username != "user1" {
		t.Fatalf("expected username 'user1', got %q", member.Username)
	}
}

func TestEditProjectMember(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/members/10", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":10,"username":"user1","access_level":40}`)
	})

	opts := &EditProjectMemberOptions{
		AccessLevel: Ptr(MasterPermission),
	}
	member, _, err := client.Projects.EditProjectMember(context.Background(), 1, 10, opts)
	if err != nil {
		t.Fatal(err)
	}
	if member.AccessLevel != MasterPermission {
		t.Fatalf("expected access level %d, got %d", MasterPermission, member.AccessLevel)
	}
}

func TestDeleteProjectMember(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/1/members/10", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Projects.DeleteProjectMember(context.Background(), 1, 10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProjectExtendedOperations(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/projects/owned", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":2,"name":"owned"}]`)
	})
	mux.HandleFunc("/api/v3/projects/1/members/10", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"id":10,"username":"user1","access_level":30}`)
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"id":10,"username":"user1","access_level":40}`)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v3/projects/1", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"id":1,"name":"updated"}`)
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"id":1,"name":"my-project"}`)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v3/projects/1/share", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/api/v3/projects/1/shares", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"project_id":1,"group_id":2,"group_access":30}]`)
	})
	mux.HandleFunc("/api/v3/projects/1/share/2", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/api/v3/projects/1/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"project_id":1,"action_name":"CREATED"}]`)
	})
	mux.HandleFunc("/api/v3/projects/1/star", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPut, http.MethodGet:
			fmt.Fprint(w, `true`)
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v3/projects/1/stars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"project_id":1,"user":{"id":10,"username":"user1"}}]`)
	})

	member, _, err := client.Projects.GetProjectMember(context.Background(), 1, 10)
	if err != nil || member.ID != 10 {
		t.Fatalf("get project member failed: %+v, %v", member, err)
	}
	project, _, err := client.Projects.UpdateProject(context.Background(), 1, &UpdateProjectOptions{Name: Ptr("updated")})
	if err != nil || project.Name != "updated" {
		t.Fatalf("update project failed: %+v, %v", project, err)
	}
	owned, _, err := client.Projects.ListOwnedProjects(context.Background(), nil)
	if err != nil || len(owned) != 1 {
		t.Fatalf("list owned projects failed: %+v, %v", owned, err)
	}
	if _, err := client.Projects.ShareProject(context.Background(), 1, &ShareProjectOptions{GroupID: Ptr(2), GroupAccess: Ptr(30)}); err != nil {
		t.Fatal(err)
	}
	shares, _, err := client.Projects.ListProjectShares(context.Background(), 1)
	if err != nil || len(shares) != 1 {
		t.Fatalf("list shares failed: %+v, %v", shares, err)
	}
	if _, err := client.Projects.DeleteProjectShare(context.Background(), 1, 2); err != nil {
		t.Fatal(err)
	}
	events, _, err := client.Projects.ListProjectEvents(context.Background(), 1, nil)
	if err != nil || len(events) != 1 {
		t.Fatalf("list events failed: %+v, %v", events, err)
	}
	starred, _, err := client.Projects.StarProject(context.Background(), 1)
	if err != nil || !starred {
		t.Fatalf("star project failed: %v, %v", starred, err)
	}
	starStatus, _, err := client.Projects.GetStarStatus(context.Background(), 1)
	if err != nil || !starStatus {
		t.Fatalf("get star status failed: %v, %v", starStatus, err)
	}
	stars, _, err := client.Projects.ListProjectStars(context.Background(), 1, nil)
	if err != nil || len(stars) != 1 {
		t.Fatalf("list project stars failed: %+v, %v", stars, err)
	}
	if _, err := client.Projects.UnstarProject(context.Background(), 1); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Projects.DeleteProject(context.Background(), 1); err != nil {
		t.Fatal(err)
	}
}
