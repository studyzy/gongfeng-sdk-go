package gongfeng

import (
	"context"
	"fmt"
	"net/http"
)

// SSHKey 表示用户 SSH Key。
type SSHKey struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Key       string `json:"key"`
	CreatedAt Time   `json:"created_at"`
}

// Email 表示用户邮箱。
type Email struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// UsersService 处理与用户相关的 API 调用。
type UsersService struct {
	client *Client
}

// ListUsersOptions 是 ListUsers 的可选参数。
type ListUsersOptions struct {
	ListOptions
	Search *string `url:"search,omitempty" json:"search,omitempty"`
}

// ListUsers 获取用户列表。
func (s *UsersService) ListUsers(ctx context.Context, opts *ListUsersOptions) ([]*User, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "users", opts)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// ListWatchedProjectsOptions 是 ListWatchedProjects 的可选参数。
type ListWatchedProjectsOptions struct {
	ListOptions
}

// ListWatchedProjects 获取当前用户关注的项目列表。
func (s *UsersService) ListWatchedProjects(ctx context.Context, opts *ListWatchedProjectsOptions) ([]*Project, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "user/watched", opts)
	if err != nil {
		return nil, nil, err
	}

	var projects []*Project
	resp, err := s.client.Do(req, &projects)
	if err != nil {
		return nil, resp, err
	}

	return projects, resp, nil
}

// GetUser 获取指定 ID 或用户名的用户信息。
func (s *UsersService) GetUser(ctx context.Context, uid interface{}) (*UserDetail, *Response, error) {
	userID, err := parseID(uid)
	if err != nil {
		return nil, nil, err
	}
	path := fmt.Sprintf("users/%s", userID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var user UserDetail
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}

// CreateSSHKeyOptions 是 CreateSSHKey 的可选参数。
type CreateSSHKeyOptions struct {
	Title *string `json:"title,omitempty" url:"title,omitempty"`
	Key   *string `json:"key,omitempty" url:"key,omitempty"`
}

// CreateSSHKey 为当前用户创建 SSH Key。
func (s *UsersService) CreateSSHKey(ctx context.Context, opts *CreateSSHKeyOptions) (*SSHKey, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "user/keys", opts)
	if err != nil {
		return nil, nil, err
	}

	var key SSHKey
	resp, err := s.client.Do(req, &key)
	if err != nil {
		return nil, resp, err
	}

	return &key, resp, nil
}

// ListSSHKeys 获取当前用户的 SSH Key 列表。
func (s *UsersService) ListSSHKeys(ctx context.Context) ([]*SSHKey, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "user/keys", nil)
	if err != nil {
		return nil, nil, err
	}

	var keys []*SSHKey
	resp, err := s.client.Do(req, &keys)
	if err != nil {
		return nil, resp, err
	}

	return keys, resp, nil
}

// GetSSHKey 获取当前用户指定的 SSH Key。
func (s *UsersService) GetSSHKey(ctx context.Context, keyID int) (*SSHKey, *Response, error) {
	path := fmt.Sprintf("user/keys/%d", keyID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var key SSHKey
	resp, err := s.client.Do(req, &key)
	if err != nil {
		return nil, resp, err
	}

	return &key, resp, nil
}

// DeleteSSHKey 删除当前用户指定的 SSH Key。
func (s *UsersService) DeleteSSHKey(ctx context.Context, keyID int) (*Response, error) {
	path := fmt.Sprintf("user/keys/%d", keyID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// CreateEmailOptions 是 CreateEmail 的可选参数。
type CreateEmailOptions struct {
	Email *string `json:"email,omitempty" url:"email,omitempty"`
}

// CreateEmail 为当前用户添加邮箱。
func (s *UsersService) CreateEmail(ctx context.Context, opts *CreateEmailOptions) (*Email, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "user/emails", opts)
	if err != nil {
		return nil, nil, err
	}

	var email Email
	resp, err := s.client.Do(req, &email)
	if err != nil {
		return nil, resp, err
	}

	return &email, resp, nil
}

// GetUserByEmailOptions 是 GetUserByEmail 的可选参数。
type GetUserByEmailOptions struct {
	Email *string `url:"email,omitempty" json:"email,omitempty"`
}

// GetUserByEmail 通过邮箱获取用户基本信息。
func (s *UsersService) GetUserByEmail(ctx context.Context, opts *GetUserByEmailOptions) (*User, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "user/email", opts)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}

// ListEmails 获取当前用户邮箱列表。
func (s *UsersService) ListEmails(ctx context.Context) ([]*Email, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "user/emails", nil)
	if err != nil {
		return nil, nil, err
	}

	var emails []*Email
	resp, err := s.client.Do(req, &emails)
	if err != nil {
		return nil, resp, err
	}

	return emails, resp, nil
}

// GetEmail 获取当前用户指定邮箱。
func (s *UsersService) GetEmail(ctx context.Context, emailID int) (*Email, *Response, error) {
	path := fmt.Sprintf("user/emails/%d", emailID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var email Email
	resp, err := s.client.Do(req, &email)
	if err != nil {
		return nil, resp, err
	}

	return &email, resp, nil
}

// DeleteEmail 删除当前用户指定邮箱。
func (s *UsersService) DeleteEmail(ctx context.Context, emailID int) (*Response, error) {
	path := fmt.Sprintf("user/emails/%d", emailID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// GetCurrentUser 获取当前认证用户的信息。
func (s *UsersService) GetCurrentUser(ctx context.Context) (*UserDetail, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "user", nil)
	if err != nil {
		return nil, nil, err
	}

	var user UserDetail
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}
