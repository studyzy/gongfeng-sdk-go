package gongfeng

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListUsers(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("PRIVATE-TOKEN") == "" {
			t.Fatal("missing PRIVATE-TOKEN header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Total", "1")
		fmt.Fprint(w, `[{"id":1,"username":"admin","name":"Admin"}]`)
	})

	users, resp, err := client.Users.ListUsers(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}
	if users[0].Username != "admin" {
		t.Fatalf("expected username 'admin', got %q", users[0].Username)
	}
	if users[0].Name != "Admin" {
		t.Fatalf("expected name 'Admin', got %q", users[0].Name)
	}
	if resp.TotalItems != 1 {
		t.Fatalf("expected TotalItems=1, got %d", resp.TotalItems)
	}
}

func TestGetUser(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/users/1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"username":"admin","name":"Admin","email":"admin@test.com"}`)
	})

	user, _, err := client.Users.GetUser(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID != 1 {
		t.Fatalf("expected ID=1, got %d", user.ID)
	}
	if user.Username != "admin" {
		t.Fatalf("expected username 'admin', got %q", user.Username)
	}
	if user.Email != "admin@test.com" {
		t.Fatalf("expected email 'admin@test.com', got %q", user.Email)
	}
}

func TestGetCurrentUser(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":1,"username":"me","name":"Me","email":"me@test.com"}`)
	})

	user, _, err := client.Users.GetCurrentUser(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if user.Username != "me" {
		t.Fatalf("expected username 'me', got %q", user.Username)
	}
	if user.Email != "me@test.com" {
		t.Fatalf("expected email 'me@test.com', got %q", user.Email)
	}
}

func TestListWatchedProjects(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/user/watched", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"name":"watched"}]`)
	})

	projects, _, err := client.Users.ListWatchedProjects(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) != 1 || projects[0].Name != "watched" {
		t.Fatalf("unexpected projects: %+v", projects)
	}
}

func TestSSHKeyOperations(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/user/keys", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPost:
			fmt.Fprint(w, `{"id":1,"title":"k","key":"ssh-rsa xxx"}`)
		case http.MethodGet:
			fmt.Fprint(w, `[{"id":1,"title":"k","key":"ssh-rsa xxx"}]`)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v3/user/keys/1", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"id":1,"title":"k","key":"ssh-rsa xxx"}`)
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})

	key, _, err := client.Users.CreateSSHKey(context.Background(), &CreateSSHKeyOptions{Title: Ptr("k"), Key: Ptr("ssh-rsa xxx")})
	if err != nil || key.ID != 1 {
		t.Fatalf("create key failed: %+v, %v", key, err)
	}
	keys, _, err := client.Users.ListSSHKeys(context.Background())
	if err != nil || len(keys) != 1 {
		t.Fatalf("list keys failed: %+v, %v", keys, err)
	}
	gotKey, _, err := client.Users.GetSSHKey(context.Background(), 1)
	if err != nil || gotKey.ID != 1 {
		t.Fatalf("get key failed: %+v, %v", gotKey, err)
	}
	if _, err := client.Users.DeleteSSHKey(context.Background(), 1); err != nil {
		t.Fatal(err)
	}
}

func TestEmailOperations(t *testing.T) {
	client, mux := setup(t)

	mux.HandleFunc("/api/v3/user/emails", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPost:
			fmt.Fprint(w, `{"id":1,"email":"a@test.com"}`)
		case http.MethodGet:
			fmt.Fprint(w, `[{"id":1,"email":"a@test.com"}]`)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})
	mux.HandleFunc("/api/v3/user/email", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"username":"user1","name":"User 1"}`)
	})
	mux.HandleFunc("/api/v3/user/emails/1", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"id":1,"email":"a@test.com"}`)
		case http.MethodDelete:
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	})

	email, _, err := client.Users.CreateEmail(context.Background(), &CreateEmailOptions{Email: Ptr("a@test.com")})
	if err != nil || email.ID != 1 {
		t.Fatalf("create email failed: %+v, %v", email, err)
	}
	user, _, err := client.Users.GetUserByEmail(context.Background(), &GetUserByEmailOptions{Email: Ptr("a@test.com")})
	if err != nil || user.Username != "user1" {
		t.Fatalf("get user by email failed: %+v, %v", user, err)
	}
	emails, _, err := client.Users.ListEmails(context.Background())
	if err != nil || len(emails) != 1 {
		t.Fatalf("list emails failed: %+v, %v", emails, err)
	}
	gotEmail, _, err := client.Users.GetEmail(context.Background(), 1)
	if err != nil || gotEmail.ID != 1 {
		t.Fatalf("get email failed: %+v, %v", gotEmail, err)
	}
	if _, err := client.Users.DeleteEmail(context.Background(), 1); err != nil {
		t.Fatal(err)
	}
}
