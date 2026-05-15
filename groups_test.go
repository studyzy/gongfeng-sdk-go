package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"name":"my-group","path":"my-group"}`)
	})

	opts := &CreateGroupOptions{
		Name: Ptr("my-group"),
		Path: Ptr("my-group"),
	}
	group, _, err := client.Groups.CreateGroup(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if group.ID != 1 {
		t.Fatalf("expected ID=1, got %d", group.ID)
	}
	if group.Name != "my-group" {
		t.Fatalf("expected name 'my-group', got %q", group.Name)
	}
}

func TestEditGroup(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"name":"updated-group","path":"my-group"}`)
	})

	opts := &EditGroupOptions{
		Name: Ptr("updated-group"),
	}
	group, _, err := client.Groups.EditGroup(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if group.Name != "updated-group" {
		t.Fatalf("expected name 'updated-group', got %q", group.Name)
	}
}

func TestDeleteGroup(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Groups.DeleteGroup(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListGroups(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"name":"group1"}]`)
	})

	groups, _, err := client.Groups.ListGroups(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(groups))
	}
	if groups[0].Name != "group1" {
		t.Fatalf("expected name 'group1', got %q", groups[0].Name)
	}
}

func TestGetGroup(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"name":"my-group","path":"my-group"}`)
	})

	group, _, err := client.Groups.GetGroup(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if group.ID != 1 {
		t.Fatalf("expected ID=1, got %d", group.ID)
	}
	if group.Path != "my-group" {
		t.Fatalf("expected path 'my-group', got %q", group.Path)
	}
}

func TestListGroupMembers(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups/1/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":10,"username":"user1","access_level":30}]`)
	})

	members, _, err := client.Groups.ListGroupMembers(context.Background(), 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 1 {
		t.Fatalf("expected 1 member, got %d", len(members))
	}
	if members[0].Username != "user1" {
		t.Fatalf("expected username 'user1', got %q", members[0].Username)
	}
	if members[0].AccessLevel != DeveloperPermission {
		t.Fatalf("expected access_level=%d, got %d", DeveloperPermission, members[0].AccessLevel)
	}
}

func TestAddGroupMember(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups/1/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":10,"username":"user1","access_level":30}`)
	})

	opts := &AddGroupMemberOptions{
		UserID:      Ptr(10),
		AccessLevel: Ptr(DeveloperPermission),
	}
	member, _, err := client.Groups.AddGroupMember(context.Background(), 1, opts)
	if err != nil {
		t.Fatal(err)
	}
	if member.ID != 10 {
		t.Fatalf("expected ID=10, got %d", member.ID)
	}
}

func TestEditGroupMember(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups/1/members/10", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":10,"username":"user1","access_level":40}`)
	})

	opts := &EditGroupMemberOptions{
		AccessLevel: Ptr(MasterPermission),
	}
	member, _, err := client.Groups.EditGroupMember(context.Background(), 1, 10, opts)
	if err != nil {
		t.Fatal(err)
	}
	if member.AccessLevel != MasterPermission {
		t.Fatalf("expected access_level=%d, got %d", MasterPermission, member.AccessLevel)
	}
}

func TestDeleteGroupMember(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/groups/1/members/10", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	})

	_, err := client.Groups.DeleteGroupMember(context.Background(), 1, 10)
	if err != nil {
		t.Fatal(err)
	}
}
